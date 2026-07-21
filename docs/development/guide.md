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

## Pull Request Workflow
1. Develop on a feature branch (`feature/your-feature`).
2. Run `make test` locally.
3. Open a Pull Request against `main`.
4. Wait for CI checks to pass.
5. Request review from maintainers.

