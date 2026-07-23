# Engineering Lifecycle

This document defines the macro-level engineering sprint rhythm for the repository.

## The Engineering Loop
All feature work must adhere strictly to the following lifecycle:

1. **Design Review:** All architecture and interfaces must be designed and approved before implementation.
2. **Implementation:** A single, limited-scope sprint executing the approved design.
3. **Local Verification:** Formatting (gofmt/gofumpt), linting (golangci-lint), and test coverage checks run locally.
4. **Code Review:** Ensure no architectural leakage or design deviations occurred.
5. **CI Verification:** Push to trigger GitHub Actions; pipeline must be 100% Green.
6. **Merge:** Integrate feature into the main branch.

## Release Gate Rule (Non-Negotiable)
- No feature work may begin or continue while the default branch (`main`) has failing required CI checks.
- Every release and merge must leave `main` in a fully green state before the next engineering sprint begins.

## Branching & Commits
- Use `feature/*` for new features, `fix/*` for bugfixes, and `chore/*` for maintenance.
- Commit messages must follow Conventional Commits format (e.g., `feat(core): ...`, `fix(ci): ...`).
