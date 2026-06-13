<!-- file: .github/copilot-instructions.md -->
<!-- version: 1.0.0 -->
<!-- guid: a3f7e2d1-8b4c-4e6a-9f1d-2c5b3a7e8f0d -->
<!-- last-edited: 2026-06-13 -->

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
