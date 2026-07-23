# Developer Guide

Welcome to the Base Infrastructure project. This guide outlines everything you need to start contributing to the codebase.

## Setup
Ensure you have installed Go 1.21 or higher.

```bash
git clone https://github.com/Hardik-Sankhla/base-infrastructure.git
cd base-infrastructure
make
```

## Coding Standards
- **Formatting**: The project enforces strict `gofmt` and `gofumpt` checks in CI. Run `make lint` before pushing.
- **Errors**: Always wrap errors. Use `fmt.Errorf("failed to do X: %w", err)`.
- **Comments**: All exported functions, methods, and structs must have a descriptive comment block.

## Commit Convention
We strictly adhere to Conventional Commits:
- `feat:` for new features (e.g. `feat: implement OS discovery stage`).
- `fix:` for bug fixes.
- `chore:` for maintenance (e.g. `chore: update dependencies`).
- `docs:` for documentation updates.
- `test:` for adding missing tests.

## Pull Request Workflow (Engineering OS)
This repository is engineered by AI agents using the **Engineering OS v2.0** (`.ai/`). Human contributors must follow the same lifecycle:

1. **Design Review:** All architecture must be designed and approved by a maintainer acting as the "Architect".
2. **Implementation:** Develop on a feature branch (`feature/your-feature`), following the strict boundaries in `.ai/ARCHITECTURE_RULES.md`.
3. **Verification:** Run all local tests and linters (`make test`, `make lint`).
4. **CI Checks:** Open a Pull Request and wait for CI checks to pass. No code is merged if CI is red.
5. **Review:** Request review from human maintainers acting as the "Reviewer".

