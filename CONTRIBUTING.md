# Contributing to Base Infrastructure

First off, thank you for considering contributing to Base Infrastructure! 

This repository is primarily engineered by AI agents using the **Engineering OS v2.0** located in the `.ai/` directory. Human contributors must interact with the repository within these boundaries.

## Development Guide
For comprehensive instructions on how to set up your environment, follow our [Developer Guide](docs/development/guide.md).

## Code of Conduct
By participating, you are expected to uphold our [Code of Conduct](CODE_OF_CONDUCT.md).

## Contributor Workflow (Engineering OS)
Whether you are writing code yourself or delegating to an AI agent, you must follow the `.ai/ENGINEERING_LIFECYCLE.md`:

1. **Design Review:** All architecture must be designed and approved (`.ai/checklists/design_review.md`).
2. **Architecture Approval:** Human maintainers act as the "Architect" persona to approve designs before implementation begins.
3. **Implementation:** Code is written on a feature branch (`feature/your-feature-name`), following the `.ai/checklists/implementation.md`.
4. **Verification:** Run all local tests (`make test`, `make verify`). Ensure code passes `.ai/checklists/verification.md`.
5. **CI Green:** Ensure GitHub Actions are 100% Green. No feature work may continue if `main` is broken.
6. **Merge:** Human maintainers act as the "Reviewer" to enforce architectural boundaries and merge code.
7. **Release:** Use `.ai/checklists/release.md` to safely cut a new tagged release.

Please follow our Conventional Commits rules detailed in the Developer Guide.
