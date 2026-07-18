<!-- file: docs/AUDIT-2026-07-18-post-org-migration.md -->
<!-- version: 1.0.0 -->
<!-- guid: 543a463e-c1e3-45c9-8362-a0230354c7c5 -->
<!-- last-edited: 2026-07-18 -->

# falkcorp/gcommon — Post-Org-Migration Audit (2026-07-18)

## Executive Summary

The repo builds and vets clean, which is exactly why this went unnoticed for as long as it has. In reality, **every single generated protobuf package in the repo is broken at runtime** — commit `08c26f44`'s jdfalk→falkcorp org-rename corrupted the embedded `rawDesc` FileDescriptor bytes in all ~3521–3523 `.pb.go` files across all 10 `pkg/*pb` modules (v1 and v2), and three independent auditors reproduced the same `panic: runtime error: slice bounds out of range` at package `init()` in all 20 `pkg/*pb[/v2]` packages. `go build`/`go vet` never execute `init()` so they report 0 failures everywhere; the repo has zero test files in `pkg/`, so `go test` is equally blind — except in `services/health`, whose real (non-scratch) test suite is the one and only genuine, currently-red CI-visible signal of this bug on HEAD. Layered on top: `buf.gen.yaml`'s managed-mode overrides are independently broken in two ways (stale path patterns, wrong prefix values) so regeneration today would silently keep falling back to the stale `jdfalk` go_package; 8 of 10 `pkg/*pb` families can't even be `go get`-resolved externally due to stale cross-module version pins masked by local `replace` directives; all 89 existing git tags are unreachable under Go's nested-module tag-naming convention and predate both the directory rename and the org transfer; and CI on `main` is not merely blind to this — it is completely non-functional (0 jobs run, every recent `ci.yml` run fails before checkout) because the reusable workflow it calls is pinned to an orphaned, unreachable commit SHA in `falkcorp/github-common`. Org-migration string-replacement itself is nearly done (only ~13 stale `jdfalk` references remain, all in low-traffic tooling/docs, not in go.mod/buf.yaml/README). The gRPC service-completeness audit did not return results and that area remains unverified. Net assessment: this repo is **not safe to consume or release in its current state**, and the path to green requires full protobuf regeneration, a hand-fix of every stale go.mod pin, a fresh tag ladder, and a CI repair — in that rough order.

## Critical Blockers

| # | Severity | Blocker | Found by |
|---|---|---|---|
| 1 | Critical | All 20 `pkg/*pb[/v2]` packages panic at `init()` (`slice bounds out of range`) due to corrupted `rawDesc` FileDescriptor bytes from commit `08c26f44` — 100% of generated code, confirmed via 3 independent methods | Protobuf audit, Module/Tagging audit, Build/Vet/Test audit (3-way confirmed) |
| 2 | Critical | `go build`/`go vet` are 0% effective at catching the above (never run `init()`); zero test files exist under `pkg/` | Build/Vet/Test audit |
| 3 | Critical | `services/health`'s real test suite is already red on HEAD from this exact bug — the only genuine CI-visible signal in the repo | Build/Vet/Test audit |
| 4 | Critical | CI (`ci.yml`) is completely non-functional on `main` — 0 jobs run on every recent push/PR because the reusable workflow is pinned to an orphaned/unreachable commit (`83352d45a9…`) in `falkcorp/github-common` | Org-migration audit |
| 5 | Critical | `buf.gen.yaml`'s managed-mode `path:` overrides (all 20 entries) never match anything in the real BSR layout (missing `pb` suffix), so regeneration today falls through to the still-stale `jdfalk` go_package in BSR source | Protobuf audit |
| 6 | High | `buf.gen.yaml`'s `value:` overrides are also wrong (full package path instead of bare prefix) — a second, independent bug that must be fixed together with #5 | Protobuf audit |
| 7 | High | 8 of 10 `pkg/*pb` families (v1 and v2) are unreachable via `go get` for any external consumer — stale pre-migration `commonpb` pseudo-version pins, masked locally by `replace` directives | Module/Tagging audit |
| 8 | High | All 89 existing git tags are permanently unreachable under Go's nested-module tag convention (wrong names, predate the `pkg/common`→`pkg/commonpb` rename and org transfer) | Module/Tagging audit |
| 9 | High | Root module has no usable tag or pseudo-version at all — no `go.mod`/`main.go` at the latest tag (`v2.1.6`), and its own pseudo-version path also fails via the same stale-pin bug | Module/Tagging audit |
| 10 | High | `services/auth`'s green test suite is a false positive — every `authpb`-consuming method is inside a `/* TODO: uncomment when protobuf issues are resolved */` block comment; no real gRPC/protobuf functionality is wired up | Build/Vet/Test audit |
| 11 | High | Structural mismatch between BSR's symmetric `v1/`+`v2/` proto layout and the repo's asymmetric on-disk Go convention (unsuffixed `pkg/commonpb` = v1, `pkg/commonpb/v2`) — buf's `go_package_prefix` mechanism cannot reproduce the current layout; this is a design decision, not a config fix | Protobuf audit |
| 12 | Medium | `.goreleaser.yml`'s `release.github.owner` is still `jdfalk`; `scripts/manage_releases.py` defaults to `jdfalk/gcommon` and has a hardcoded now-nonexistent path | Org-migration audit, Module/Tagging audit |

## Findings by Area

### Protobuf Code Generation

- **Universal descriptor corruption, double-desync.** A byte-level parser run against all 3521 `.pb.go` files touched by commit `08c26f44` found every one corrupted identically: the length-prefix byte for the `go_package` field (protobuf tag `0x5A`) was left at the old (`jdfalk`) string length after the string itself was rewritten to the 2-bytes-longer `falkcorp` string. Worse, the *enclosing* `FileOptions` submessage length prefix (tag `0x42`) is also stale by the same +2 offset — a compounding double desync, which is why the failure mode is a hard panic rather than a merely-wrong string. Example: `pkg/commonpb/v2/ack_level.pb.go` declares length `0x29`=41 for `go_package` but the actual `falkcorp` string is 43 bytes; line 91 still reads the outer length as `\x96\x01` (150) though it should be 152.
- Breakdown by module (all 100% corrupted): authpb 53/53, commonpb 1018/1018, configpb 228/228, databasepb 278/278, healthpb 72/72, mediapb 150/150, metricspb 440/440, organizationpb 234/234, queuepb 644/644, webpb 404/404.
- **`buf.gen.yaml` cannot regenerate correctly today, for two independent reasons.** (1) All 20 override `path:` entries reference a stale layout missing the `pb` suffix (`path: common/v1` instead of `commonpb/v1`), so none of them ever match a real BSR file — confirmed via `buf ls-files buf.build/falkcorp/gcommon` (3521 files, all under `<name>pb/v{1,2}/`). With zero overrides firing, generation falls through to the BSR source's own hardcoded `option go_package = "github.com/jdfalk/gcommon/pkg/commonpb/v2";` (confirmed stale in `commonpb`, `databasepb`, `queuepb`, `metricspb` samples). (2) Even with the path bug fixed, the `value:` fields are the *full* target package (`github.com/falkcorp/gcommon/pkg/commonpb/v2`) rather than a bare prefix; `go_package_prefix` concatenates automatically, so this produces a doubled path (`.../pkg/commonpb/v2/commonpb/v2;commonv2`). Both bugs were proven experimentally and fully reverted (`git checkout -- buf.gen.yaml`, confirmed clean).
- **v1/v2 directory-layout mismatch.** BSR always splits `v1/`/`v2/` into explicit sibling proto directories; the checked-in Go tree uses an asymmetric convention (unsuffixed `pkg/commonpb` = v1, `pkg/commonpb/v2`). A corrected `buf generate` run writes v1 output to a *new* `pkg/commonpb/v1/` directory rather than the existing unsuffixed one — this is what produces the untracked `pkg/*pb/v1` directories seen after `buf generate`. This needs an explicit human decision (see Open Questions), not a silent fix.
- **`go build`/`go vet` are structurally blind.** A throwaway module importing `pkg/commonpb/v2` via a local `replace` panics on `go run`; `go build ./...` and `go vet ./...` both exit 0 in the same package because neither executes `init()`. Zero `*_test.go` files exist anywhere under `pkg/`.
- `buf.build/falkcorp/gcommon` is not pinned in `buf.lock` (only `bufbuild/protovalidate` and `googleapis/googleapis` are), so every `buf generate` run is non-reproducible.
- Once both `buf.gen.yaml` bugs are fixed, regeneration correctly produces falkcorp-orged, byte-correct code *without* needing to touch BSR source — proven experimentally — though the BSR source itself is still separately stale and will bite any other consumer generating without an equivalent override.

### gRPC Service Completeness

**This audit returned no data** (`null` — the specialist run either failed or produced an empty result; it was not merely "no findings"). No service-completeness matrix is available for this report. Given the confirmed protobuf corruption, treat gRPC service behavior as **entirely unverified** for this report — do not infer "clean" from its absence. This should be re-run before the repo is considered release-ready; see Open Questions.

### Module Structure & Release Tagging

Module structure itself is sound: exactly 24 `go.mod` files, every module path self-consistent with its directory, every cross-module `require` has a matching local `replace`, and `go build ./...` succeeds in all 24 directories via local replaces alone (the previously-tracked "9 submodules missing replace directives" issue is fully closed). Everything below the local-build layer, however, is broken:

| Module | Directory | Self-consistent | `@latest` tag resolves | Pseudo-version resolves | Notes |
|---|---|---|---|---|---|
| `github.com/falkcorp/gcommon` | `.` | ✅ | ✅ (v2.1.6) | ❌ | v2.1.6 predates `go.mod`/`main.go` entirely; `go install @v2.1.6` fails "does not contain package". HEAD pseudo-version also fails (stale `internal` pin). |
| `.../internal` | `internal/` | ✅ | ❌ | ❌ | No tags exist; also unimportable externally by Go's `internal/` rule regardless. |
| `.../pkg/authpb` | `pkg/authpb/` | ✅ | ❌ | ✅ | Zero tags at all (only pb family with none). Leaf module, pseudo-version resolves but panics at `init()`. |
| `.../pkg/authpb/v2` | `pkg/authpb/v2/` | ✅ | ❌ | ✅ | Same as v1; confirmed panic via built+run binary. |
| `.../pkg/commonpb` | `pkg/commonpb/` | ✅ | ❌ | ✅ | Tag `pkg/common/v2.1.5` exists but declares itself `module github.com/jdfalk/gcommon/pkg/common` — unreachable regardless. Pseudo-version resolves, panics at runtime. |
| `.../pkg/commonpb/v2` | `pkg/commonpb/v2/` | ✅ | ❌ | ✅ | `@v2.1.5` falls through to fetching the *root* repo's tag and fails "does not contain package". Pseudo-version resolves, panics (original confirmed example). |
| `.../pkg/configpb` | `pkg/configpb/` | ✅ | ❌ | ❌ | Same stale-pin pattern as v2 sibling (not directly re-tested; structurally identical require line). |
| `.../pkg/configpb/v2` | `pkg/configpb/v2/` | ✅ | ❌ | ❌ | Directly tested: fails with commonpb/v2 path-mismatch error. |
| `.../pkg/databasepb` | `pkg/databasepb/` | ✅ | ❌ | ❌ | Directly tested: same failure (v1 track). |
| `.../pkg/databasepb/v2` | `pkg/databasepb/v2/` | ✅ | ❌ | ❌ | Directly tested. Workaround (co-pinning `commonpb/v2` to same SHA) resolves+builds but panics at runtime. |
| `.../pkg/healthpb` | `pkg/healthpb/` | ✅ | ❌ | ❌ | Same stale-pin pattern (not re-tested). |
| `.../pkg/healthpb/v2` | `pkg/healthpb/v2/` | ✅ | ❌ | ❌ | Directly tested: fails. |
| `.../pkg/mediapb` | `pkg/mediapb/` | ✅ | ❌ | ❌ | Same pattern (not re-tested). |
| `.../pkg/mediapb/v2` | `pkg/mediapb/v2/` | ✅ | ❌ | ❌ | Directly tested: fails. |
| `.../pkg/metricspb` | `pkg/metricspb/` | ✅ | ❌ | ❌ | Same pattern (not re-tested). |
| `.../pkg/metricspb/v2` | `pkg/metricspb/v2/` | ✅ | ❌ | ❌ | Directly tested: fails. |
| `.../pkg/organizationpb` | `pkg/organizationpb/` | ✅ | ❌ | ❌ | Same pattern (not re-tested). |
| `.../pkg/organizationpb/v2` | `pkg/organizationpb/v2/` | ✅ | ❌ | ❌ | Directly tested: fails. |
| `.../pkg/queuepb` | `pkg/queuepb/` | ✅ | ❌ | ❌ | Same pattern (not re-tested). |
| `.../pkg/queuepb/v2` | `pkg/queuepb/v2/` | ✅ | ❌ | ❌ | Directly tested: fails. |
| `.../pkg/webpb` | `pkg/webpb/` | ✅ | ❌ | ❌ | Same pattern (not re-tested). |
| `.../pkg/webpb/v2` | `pkg/webpb/v2/` | ✅ | ❌ | ❌ | Directly tested: fails. |
| `.../services` | `services/` | ✅ | ❌ | ❌ | No tags. Same stale-pin pattern for `commonpb`/`healthpb` requires. |
| `.../services/auth` | `services/auth/` | ✅ | ❌ | ❌ | No tags. `pkg/authpb` pinned at Go's zero-placeholder version, resolvable only via local replace. |

Additional findings in this area:
- **Tag namespace is dead.** All 89 existing tags across root + 9 pb families (not `authpb`, which has zero) share a single creation date (2025-09-26), use the pre-rename `pkg/common` directory name, and don't follow Go's `<module-subdir>/vX.Y.Z` nested-tag convention — `go get` silently falls through to fetching the unrelated root tree instead of erroring clearly.
- **Root's semver is `+incompatible`.** `go.mod` declares `module github.com/falkcorp/gcommon` with no `/v2` suffix despite v2.x tags, so every v2.x tag resolves as `+incompatible` and loses Go's major-version guarantees. Lower-impact since root is `package main`, but should be fixed or explicitly documented.
- **Release tooling is broken and stale.** `scripts/manage_releases.py` defaults `repo_owner="jdfalk"`, `--repo` defaults to `jdfalk/gcommon`, and hardcodes `cwd="/Users/jdfalk/repos/github.com/jdfalk/gcommon"` (line 113) — a path that no longer exists. `.goreleaser.yml` line 49: `owner: jdfalk`. The script actually wired to `make release-*` (`scripts/release-manager.py`) only ever cuts a single root `vX.Y.Z` tag (`version_pattern = ^v?\d+\.\d+\.\d+$`, lines 106–120) with no logic for the `pkg/*pb` nested-module ladder at all, and its go-mod-tidy step only walks `pkg/` (lines 162–163), skipping root/internal/services/services-auth.
- Minor: root `go.mod` has a dangling `replace github.com/falkcorp/gcommon/services/health => ./services/health` — `services/health` has no `go.mod` of its own and is never required, so this is inert but indicates copy-paste drift.

### Build / Vet / Test Ground Truth

Mechanical sweep at commit `94cd87ab` across all 24 `go.mod` directories (build/vet/test summary text says "23," but the reported per-module results list 24 entries — root, `internal`, 20 `pkg/*pb[/v2]`, `services`, `services/auth`; treat as 24, consistent with the Module/Tagging audit's independent count):

| Module | build | vet | test | Notes |
|---|---|---|---|---|
| `.` | pass | pass | pass | `[no test files]`; `main.go` imports no pb package. |
| `internal` | pass | pass | pass | 4 subpackages, all `[no test files]`. |
| `pkg/authpb` | pass | pass | pass | `[no test files]`; scratch test confirms panic `[-1:]`. |
| `pkg/authpb/v2` | pass | pass | pass | `[no test files]`; scratch test confirms panic `[-1:]`. |
| `pkg/commonpb` | pass | pass | pass | `[no test files]`; this is the exact package panicking inside `services/health`'s real failure. |
| `pkg/commonpb/v2` | pass | pass | pass | `[no test files]`; scratch test confirms panic `[-4:]`. |
| `pkg/configpb` / `v2` | pass | pass | pass | Same pattern, confirmed panics `[-1:]` / `[-4:]`. |
| `pkg/databasepb` / `v2` | pass | pass | pass | Same pattern, confirmed panics `[-1:]` / `[-4:]`. |
| `pkg/healthpb` / `v2` | pass | pass | pass | Same pattern, confirmed panics `[-1:]` / `[-4:]`. |
| `pkg/mediapb` / `v2` | pass | pass | pass | Same pattern, confirmed panics `[-1:]` / `[-4:]`. |
| `pkg/metricspb` / `v2` | pass | pass | pass | Same pattern, confirmed panics `[-1:]` / `[-4:]`. |
| `pkg/organizationpb` / `v2` | pass | pass | pass | Same pattern, confirmed panics `[-1:]` / `[-4:]`. |
| `pkg/queuepb` / `v2` | pass | pass | pass | Same pattern, confirmed panics `[-1:]` / `[-4:]`. |
| `pkg/webpb` / `v2` | pass | pass | pass | Same pattern, confirmed panics `[-1:]` / `[-4:]`. |
| `services` | pass | pass | **fail** | **Real, non-scratch failure**: `services/health`'s genuine test suite panics `[-1:]` tracing through `pkg/commonpb.file_commonpb_v1_ack_level_proto_init()`. This is the one CI-visible canary of the whole-repo bug today. |
| `services/auth` | pass | pass | pass | **False positive**: none of its 6 test files import `authpb`; every `authpb`-consuming method in `service.go` (~25 references, Login/ValidateToken/AuthorizeAccess/GenerateToken/RefreshToken/RevokeToken/RegisterAuthServiceServer) is wrapped in a `/* PROTOBUF-DEPENDENT METHODS (TODO: Uncomment when protobuf issues are resolved) */` block comment starting at `service.go` line ~237. |

All 20 `pkg/*pb[/v2]` panics were independently reproduced via throwaway self-import `_test.go` scratch files (constructing a zero-value message or calling `proto.Marshal`/enum `String()`), deleted immediately after each run, with `git status`/`git diff` confirmed clean afterward. v1 (unsuffixed) packages panic with `[-1:]`; v2 packages panic with `[-4:]` (note: `pkg/authpb/v2`'s Go package name is literally `authpb`, not `v2`, unlike the other nine v2 siblings).

### Org-Migration Remnants

Core migration is essentially done: every `go.mod` (root + all submodules) already declares `falkcorp`; `buf.yaml`/`buf.gen.yaml`/`buf.gen.managed.yaml` all reference `buf.build/falkcorp/gcommon`; `README.md`'s install/import instructions, `CONTRIBUTING.md`, and both `.github/workflows/*.yml` files are clean. A repo-wide case-insensitive grep for `jdfalk` found only **13 functional stragglers** plus **4 legitimate historical/attribution occurrences** (fine to leave), across 9 files:

- `.goreleaser.yml:49` — `owner: jdfalk` (latent; no workflow currently invokes goreleaser).
- `scripts/manage_releases.py` — 7 occurrences: `repo_owner` default (line 30), `--repo` argparse default (line 301), hardcoded dead `cwd=/Users/jdfalk/repos/github.com/jdfalk/gcommon` (line 113), and mixed jdfalk/falkcorp + `pkg/common`/`pkg/commonpb` references baked into generated release-note templates (lines 162–231, e.g. `buf dep update buf.build/jdfalk/gcommon:{tag}` next to `go get github.com/falkcorp/gcommon/pkg/common@{tag}` — both wrong, in the same output block).
- `scripts/create_github_issue.py:554` — `--repo` default `jdfalk/gcommon`.
- `services/auth/SUBTITLE_MANAGER_INTEGRATION.md:30` — downstream-facing example imports `github.com/jdfalk/gcommon/services/auth` (this doc specifically targets `falkcorp/subtitle-manager`, a consumer already known to be struggling with import-path/tag mismatches per user memory).
- `.github/prompts/ai-rebase-context.template.md` (lines 14, 109) — doubly-stale URL `github.com/jdfalk/copilot-agent-util-rust/releases/latest` (org *and* repo renamed to `falkcorp/safe-ai-util`; redirect currently masks the breakage). Note: the identical class of issue was just fixed in *this* repo (`ghcommon`, commit `1e4209f`) but was missed here in `gcommon`.
- `Makefile:21` — cosmetic-only echo string referencing `buf.build/jdfalk/gcommon`; actual `buf.yaml`/`buf.gen.yaml` config is already correct.

**CI status (more severe than a leftover string):** `ci.yml` calls `uses: falkcorp/github-common/.github/workflows/reusable-ci.yml@83352d45a93951d4e490a5c310e31858a52e29ce`. That commit exists as a git object in `falkcorp/github-common` (`gh api repos/falkcorp/github-common/commits/83352d45...` succeeds) but is **not reachable from any ref** (`git/trees/<sha>` 404s) — an orphaned commit, likely from a rebase/force-push on `github-common`. Every recent `ci.yml` run (`gh run list -R falkcorp/gcommon`) shows `conclusion=failure`, ~0s duration, 0 jobs, with GitHub's own diagnostic "This run likely failed because of a workflow file issue." CI currently validates **nothing** — not build, not vet, not test — and has zero visibility into the protobuf corruption or anything else. This is separate from `sync-protos.yml` ("Generate Go Code from BSR"), which also fails but for the already-documented `buf.gen.yaml` override bugs — its own repo references are correctly on `falkcorp`.

## Cross-Cutting Observations

- **The bug had no way to be caught, and this report is the first time it's been confirmed with methodological rigor.** `go build`/`go vet` cannot execute `init()`; `pkg/` has zero test files; and CI itself doesn't run at all right now due to the orphaned reusable-workflow pin. Three separate auditors (protobuf, module/tagging, build/vet/test) independently reproduced the identical panic signature across all 20 `pkg/*pb[/v2]` packages using three different methods (byte-level descriptor parsing, standalone repro modules with real `go get`, and in-package scratch tests) — this is about as solid a "confirmed, not suspected" verdict as five independent read-only audits can produce.
- **`services/health`'s test failure is the canary that was already firing and being ignored/unseen.** It is the one piece of ground truth in the whole repo that doesn't require any auditor-constructed scratch code — a real, pre-existing test genuinely fails on HEAD today. If CI were functional, this alone would have blocked merges on `services` weeks ago. The fact that it hasn't blocked anything is itself evidence that CI's non-functionality (finding in Org-Migration audit) has been masking real signal, not just theoretical corruption.
- **`services/auth` is a second, more troubling instance of masking — this time human-authored.** Someone already hit the exact protobuf panic and, instead of filing/fixing it, commented out every protobuf-dependent method in `service.go` with a `TODO: uncomment when protobuf issues are resolved` note and left the test suite green. This means the AuthService v2 expansion (PR #2, noted in user memory as merged and "took extensive pre-existing bug fixes") is currently running with its actual gRPC surface disabled — worth flagging directly to whoever believes AuthService v2 is functional today.
- **The corruption and the `buf.gen.yaml` bugs are two independent root causes that must both be fixed, and fixing one without the other reintroduces failure from a different angle.** Naively re-running `buf generate` today (without first fixing `buf.gen.yaml`'s path/prefix bugs) would either silently regenerate stale `jdfalk`-orged code again (since BSR source itself is also stale) or, if BSR is fixed first without touching `buf.gen.yaml`, would explode the on-disk layout into new `pkg/*pb/v1/` directories per the structural-mismatch finding. The Module/Tagging audit's remediation plan and the Protobuf audit's remediation plan agree on sequencing (fix `buf.gen.yaml` → resolve v1-layout decision → regenerate → diff → test → tag), so treat them as one combined plan, not two competing ones.
- **No contradictions found between auditors on facts** — the protobuf, module/tagging, and build/vet/test audits corroborate each other's core corruption claim in ways that reinforce rather than conflict (different sampling methods, same result: 100% of `pkg/*pb[/v2]` packages). The one substantive numeric discrepancy is cosmetic: the protobuf audit counts 3521 corrupted `.pb.go` files (verified via `git show --name-only` on commit `08c26f44`), while the module/tagging audit's summary rounds to "~3523" — treat 3521 as the authoritative, directly-verified figure. The build/vet/test audit's own summary text says "23 modules" while its per-module table lists 24 — the module/tagging audit's independently-verified `find . -name go.mod | wc -l` = 24 should be treated as authoritative.
- **The gRPC service-completeness audit is a real gap in this report, not a "nothing to report" result.** Given every other audit's finding that the underlying protobuf types are unusable, it's likely any gRPC service-completeness analysis attempted to construct/inspect message or service descriptors and hit the same panic, or the audit subprocess failed outright. Either way, this needs to be explicitly re-run once the protobuf corruption is fixed — do not assume gRPC services are complete or incomplete based on this report.

## Recommended Remediation Plan

**Phase 0 — quick wins (hours, no dependencies on anything else):**
1. Re-pin `ci.yml`'s reusable-workflow reference to a commit SHA that actually resolves on `falkcorp/github-common` main (audit observed tip `0c5e8220c097fdc6dd5b8a024fa9496d761a815f` at audit time — verify current tip before using). Confirm a subsequent run actually executes jobs (`gh run list -R falkcorp/gcommon`).
2. Fix remaining `jdfalk` stragglers: `.goreleaser.yml:49` (`owner: jdfalk` → `falkcorp`), `scripts/manage_releases.py` (lines 30, 113, 162–231, 301), `scripts/create_github_issue.py:554`, `services/auth/SUBTITLE_MANAGER_INTEGRATION.md:30`, `.github/prompts/ai-rebase-context.template.md` (lines 14, 109 → `falkcorp/safe-ai-util`), `Makefile:21` (cosmetic).

**Phase 1 — fix codegen config (fast, mechanical, unblocks everything downstream):**
3. In `buf.gen.yaml`, fix all 20 override entries: correct `path:` to include the `pb` suffix (e.g. `common/v1` → `commonpb/v1`, `queue/v2` → `queuepb/v2`, …) to match the real BSR layout (verified via `buf ls-files buf.build/falkcorp/gcommon`).
4. In the same pass, fix all 20 `value:` fields to the bare shared prefix `github.com/falkcorp/gcommon/pkg` (not the full per-module path) — verify with a single-file `buf generate --path <one file>` before running repo-wide.
5. Pin `buf.gen.yaml`'s `inputs: - module: buf.build/falkcorp/gcommon` to an explicit commit or tag so future `buf generate` runs are reproducible; add the module to `buf.lock` coverage.

**Phase 2 — resolve the v1-layout decision (requires human sign-off before regenerating):**
6. Get an explicit decision on the BSR-vs-on-disk asymmetry: (A) restructure the Go tree to `pkg/<name>pb/v1/`, `pkg/<name>pb/v2/` everywhere to match BSR (breaking import-path change, but `pkg/commonpb`'s go.mod header is already marked `// Deprecated: ... Use .../pkg/commonpb/v2 instead`, softening blast radius), or (B) keep the current asymmetric convention and add a post-generate flattening step (Makefile/script) that moves v1 output up one directory and deletes the empty `v1/` dir. Recommend (A) long-term per the protobuf audit, but this must be an explicit decision, not silently defaulted.

**Phase 3 — regenerate and verify (the core fix):**
7. Run `buf generate` for the full module against the fixed config; diff file-by-file against current tree to confirm falkcorp-orged, byte-correct `go_package` strings and no unexpected structural drift beyond the org rename (BSR content may have moved since the pre-corruption baseline — diff carefully).
8. Add at least one smoke test per generated module (or one shared driver test across all `pkg/*pb` packages) that imports the package, constructs a message, and calls `proto.Marshal`/`ProtoReflect` — this is the single step that would have caught this bug immediately and prevented silent shipping.
9. Re-run the panic repro (any of the three auditors' methods) to confirm it no longer panics before committing.
10. Uncomment and re-verify the `PROTOBUF-DEPENDENT METHODS` block in `services/auth/service.go` (lines ~237–520) against the fixed `authpb` package.
11. File a follow-up (likely a different repo/pipeline) to fix the stale `option go_package = "github.com/jdfalk/...";` lines in the BSR-published proto source itself (`buf.build/falkcorp/gcommon`), so other downstream consumers without an equivalent override aren't silently affected the same way.

**Phase 4 — fix cross-module version pins (independent of regeneration, but required before any external `go get` works):**
12. Hand-edit (or temporarily remove `replace`, run `go mod tidy` against a real accessible upstream, then restore `replace`) the stale `commonpb`/`commonpb/v2` require pins in all affected `go.mod` files: `configpb[/v2]`, `databasepb[/v2]`, `healthpb[/v2]`, `mediapb[/v2]`, `metricspb[/v2]`, `organizationpb[/v2]`, `webpb[/v2]`, plus root's stale `internal` pin, `services`'s stale `commonpb`/`healthpb` pins, and `services/auth`'s zero-placeholder `authpb` pin. This cannot be caught by `go mod tidy` run inside the monorepo — local `replace` directives mask it every time.

**Phase 5 — release tooling and tagging (after phases 1–4 land):**
13. Extend or rewrite `scripts/release-manager.py` (currently only cuts a single root tag, `version_pattern = ^v?\d+\.\d+\.\d+$`) to loop over every `pkg/*pb` family and both v1/v2 directories, pushing correctly-namespaced nested-module tags (`pkg/<name>pb/vX.Y.Z`, `pkg/<name>pb/v2/vX.Y.Z`), including `authpb` (which has never had a tag). Extend its go-mod-tidy step to cover all 24 modules, not just `pkg/`.
14. Cut a fresh root tag once `go.mod`/`main.go` are reachable at HEAD; decide whether to add a `/v2` module-path suffix or explicitly document root as `go install`-only, non-semver-guaranteed.
15. Add a post-tag smoke test to the release process itself (real `go get` + build + construct-a-message round trip) — this single step would have caught both the corruption and the stale-pin defects immediately.
16. Clean up the dangling `replace github.com/falkcorp/gcommon/services/health => ./services/health` in root `go.mod` (inert, but drift).

**Phase 6 — re-verify gRPC service completeness:**
17. Re-run a gRPC service-completeness audit once phases 1–4 are done; the prior attempt returned no usable data and should not be treated as "clean."

## Open Questions

- **v1-layout decision (Phase 2, item 6 above):** is the unsuffixed `pkg/<name>pb` (v1) vs `pkg/<name>pb/v2` convention load-bearing for any current external consumer, or can it be abandoned in favor of matching BSR's symmetric layout? This determines whether remediation is a breaking rename or a permanent post-generate workaround script.
- **Who owns the BSR module itself?** The stale `option go_package = "github.com/jdfalk/...";` lines live in `buf.build/falkcorp/gcommon`'s published proto source, which is a separate artifact from this git repo (likely a different repo or a manual `buf push` pipeline) and needs its own fix path and owner.
- **Is the pre-corruption BSR commit recoverable?** `buf.lock` never pinned `buf.build/falkcorp/gcommon`, so the exact BSR state originally used to generate the currently-checked-in (pre-corruption) code may be unknown. Regeneration may therefore produce a diff larger than just the org rename — needs structural diffing, not just import-path diffing, before merging.
- **Are there existing external consumers pinned to the corrupted commit SHAs?** If so, fixing/regenerating needs a coordinated version bump communicated to them (e.g. `falkcorp/subtitle-manager`, already flagged elsewhere as struggling with gcommon import paths per prior session notes).
- **Why did the gRPC service-completeness audit return no data?** Needs to be re-run and the failure mode understood (subprocess crash vs. hit the same panic vs. timeout) before this report's gRPC section can be considered anything but "unknown."
- **Should root even be `go get`-able as a library?** It's `package main` (a server binary); its docs/usage should probably say `go install` explicitly rather than implying library consumption, independent of the `+incompatible` semver issue.
- **Is `buf.yaml`'s `deps: - buf.build/falkcorp/gcommon` intentional?** It's conceptually circular for a repo that is itself gcommon's Go-generation target — worth a sanity check on why `buf.yaml` exists here at all versus only `buf.gen.yaml`.