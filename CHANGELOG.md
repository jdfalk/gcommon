# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<!-- scriv-insert-here -->

<a id='changelog-v2.2.0'></a>
## v2.2.0 — 2026-07-19

### Changed

#### Adopt changelog fragments (`changelog.d/`) for assembling CHANGELOG.md

`CHANGELOG.md` is now assembled from per-change Markdown fragments under
`changelog.d/` by `scriv`, instead of being edited by hand. Contributors add a
fragment with `scriv create`; a CI check requires one on each PR, and the
fragments are folded into `CHANGELOG.md` when a release is published. This
removes changelog merge conflicts across parallel PRs.
