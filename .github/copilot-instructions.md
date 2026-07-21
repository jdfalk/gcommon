<!-- file: .github/copilot-instructions.md -->
<!-- version: 1.1.0 -->
<!-- guid: a3f7e2d1-8b4c-4e6a-9f1d-2c5b3a7e8f0d -->
<!-- last-edited: 2026-07-21 -->

# gcommon — Additional Context

Org-wide coding standards (file headers, language rules, commit format) are at
**https://github.com/falkcarp/.github** and apply automatically to this repo.

For full project context: **CLAUDE.md** at the repo root.

## Project overview

Go SDK for gcommon protocol buffers with shared utilities. Language: Go.

## Key directories

| Path | Purpose |
|------|---------|
| `buf.gen.yaml` | Buf code generation config |
| `buf.yaml` | Buf module definition |

## Critical constraints

- Module path uses `buf.build/falkcorp/gcommon`
- Use `git clone --recurse-submodules` to populate `.standards/`


## 📝 Changelog & TODO — Use the Fragment System (MANDATORY)

**Do not hand-edit `CHANGELOG.md`, and do not add new tasks straight into the
`TODO.md` inbox.** Both files are assembled from per-change fragments so that
parallel PRs never collide on them.

- **`CHANGELOG.md` is assembled, not hand-edited.** Add a fragment under
  `changelog.d/` (run `scriv create`, or write the Markdown file by hand). The
  fragments are folded into `CHANGELOG.md` at release time by `scriv`, and a CI
  check (`changelog-check.yml`) requires one on each PR. See `changelog.d/README.md`.
- **New `TODO.md` tasks are added via fragments.** Drop a Markdown fragment in
  `todo.d/` (see `todo.d/README.md`) instead of editing the `## 📥 Inbox`
  section. `scripts/assemble_todo.py` folds fragments in daily. This is
  **add-only**: checking a task off or removing it is a normal direct edit of
  `TODO.md`.
