# Engineering OS

Base Infrastructure is maintained and engineered using a unique, agent-driven **Engineering Operating System (Engineering OS v2.0)**. 

While `docs/` (the site you are reading now) is designed for humans, the `.ai/` directory in the repository root contains the strict rules, personas, and lifecycles used by AI agents to write code.

## The Engineering Lifecycle
Whether you are writing code yourself or delegating to an AI agent, you must follow the strict phases defined in `.ai/ENGINEERING_LIFECYCLE.md`:

1. **Design Review:** All architecture must be designed and approved (`.ai/checklists/design_review.md`).
2. **Architecture Approval:** Human maintainers act as the "Architect" persona to approve designs before implementation begins.
3. **Implementation:** Code is written on a feature branch, following the `.ai/checklists/implementation.md`.
4. **Verification:** All local tests must pass. Ensure code passes `.ai/checklists/verification.md`.
5. **CI Green:** Ensure GitHub Actions are 100% Green. No feature work may continue if `main` is broken.
6. **Merge:** Human maintainers act as the "Reviewer" to enforce boundaries and merge code.
7. **Release:** Use `.ai/checklists/release.md` to safely cut a new tagged release.

## Contribution Workflow
Human contributors should submit Pull Requests against `main`. 

1. Ensure your design is approved if you are introducing new architecture.
2. Build and verify locally: `make test` and `make lint`.
3. Submit the PR and ensure all CI Quality Gates pass.

## Release Workflow
Releases are handled by the `Release Manager` persona. When a milestone is reached, the maintainer triggers the release checklist, which generates changelogs, bumps versions, and drafts the GitHub Release.

## Repository Structure
- `.ai/` - The Agent Operating System (personas, checklists, templates)
- `cmd/platform/` - The CLI entrypoint
- `internal/` - The core application logic (bootstrap, capability builder, discovery pipeline)
- `docs/` - This human-facing documentation portal
