# Repository Map

- `internal/bootstrap/`: Dependency injection and application orchestration.
- `internal/core/`: Execution engine, pipeline definitions, and core domain interfaces.
- `internal/discovery/`: Flat capability stage implementations (`stage_*.go`).
- `internal/platform/`: OS and hardware abstraction layer (`provider_*.go`).
- `internal/services/`: Encapsulated infrastructure (e.g., PocketBase).
- `internal/presentation/`: Handles CLI formatting natively.
- `internal/capabilities/`: Builds domain structures from raw data.
- `cmd/platform/`: Core CLI execution environment (thin wrappers).
- `docs/architecture/`: Technical design docs and ADRs.
- `.github/workflows/`: The Source of Truth for validation.
