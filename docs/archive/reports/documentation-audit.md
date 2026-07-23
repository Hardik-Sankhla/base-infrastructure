# Certification: Documentation Audit

## 1. Architecture Alignment
**Correction Made:** Initially, documentation stated that Kahn's Topological Sorting algorithm was used dynamically for concurrent tier execution. 
- **Audit Finding:** A manual inspection of `internal/discovery/validator.go` and `internal/discovery/pipeline.go` proved this false. DFS is used for circular dependency validation, and Priority fields are used for standard execution sorting.
- **Resolution:** The documentation files `docs/testing/overview.md`, `docs/discovery/engine.md`, and `docs/adr/0002-pipeline-architecture.md` were updated to reflect exactly what the source code implements.

## 2. CLI Help Alignment
**Correction Made:** The previous manual stated the CLI flag was `--path` for SDK validation.
- **Audit Finding:** Executing `./platform sdk validate --path /path` natively returned `Error: unknown flag: --path`. The source code at `cmd/platform/cmd/sdk.go` explicitly leverages Cobra positional arguments (`cobra.ExactArgs(1)`).
- **Resolution:** The documentation files `docs/getting-started/quickstart.md`, `docs/sdk/overview.md`, and `docs/cli/reference.md` were corrected.

## 3. Implemented vs Planned
**Verified:** The documentation actively differentiates between the current read-only Discovery system and the planned Execution/Reconciliation system, as mandated by the project Constitution.
