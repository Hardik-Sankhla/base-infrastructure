# Release Protocol

This document governs how to transition the repository from a feature branch state to a stable, tagged release.

## 1. Definition of Release Ready
- `main` is completely stable.
- All required CI pipelines are 100% Green.
- A full Repository Hygiene Audit has been conducted (no accidental binaries, no trailing temp files).
- The `release.md` checklist is completely satisfied.

## 2. Release Steps
1. **Freeze `main`:** Ensure no lingering PRs are merged.
2. **Audit Hygiene:** Clean up generated artifacts (`deps.json`, `discovery-*.json`, etc.).
3. **Generate Release Notes:** Use the `templates/release_notes.md` standard. Summarize features, bugfixes, and breaking changes.
4. **Tag the Release:** Use Semantic Versioning (e.g., `v0.5.0`).
5. **Build & Deploy:** Wait for CI artifact deployments (GitHub Pages, Binaries).
6. **Unfreeze:** Open the next milestone branch (e.g., `feature/v0.6-xyz`).
