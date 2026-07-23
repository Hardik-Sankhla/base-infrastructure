# Architecture Freeze & Invariants

> **STATUS: FROZEN (v0.4.0)**  
> The core architecture of this repository has reached a stabilized state. Future contributors must preserve these boundaries. Structural refactoring, layer shifting, and dependency inversion are strictly prohibited without prior consensus.

## Architecture Principles

1. **One-Way Dependency Flow**: Dependencies must only flow inward toward the core domain. Outer layers may depend on inner layers; inner layers must never depend on outer layers.
2. **Interface-Driven Encapsulation**: Concrete implementations (like specific infrastructure or platform providers) must be hidden behind generic domain interfaces.
3. **Flat by Default**: Do not create deeply nested folder hierarchies for single implementations. Group cohesive files in flat packages.
4. **No God Packages**: Packages must have a single, cohesive responsibility. Orchestration and wiring belong exclusively in the `bootstrap` layer.

## Dependency Rules

The dependency flow is strictly defined as follows:
`CLI -> Bootstrap -> Core -> Platform -> Discovery -> Services`

- **CLI (`cmd/platform`)**: May only import `bootstrap`. Must not instantiate infrastructure, databases, or core engines.
- **Bootstrap (`internal/bootstrap`)**: The only package allowed to wire together `logger`, `config`, `core`, and `services`.
- **Core (`internal/core`)**: Must not import `platform`, `discovery`, `presentation`, or `services`. Contains the fundamental execution engine and domain types.
- **Discovery (`internal/discovery`)**: May import `core` and `platform`. Must not import `presentation`, `bootstrap`, or `services`.
- **Platform (`internal/platform`)**: OS-specific implementations. Must not import `discovery` or `services`.
- **Services (`internal/services`)**: Encapsulated third-party infrastructure (e.g., PocketBase). Must remain entirely isolated from the rest of the application logic.

## Package Ownership

| Package | Responsibility | Owner / Constraint |
|---------|----------------|--------------------|
| `cmd/platform` | CLI entrypoints and user commands. | Thin wrappers only. No business logic. |
| `internal/bootstrap` | Dependency injection and orchestration. | Central wiring hub. |
| `internal/core` | Pipeline execution, generic interfaces. | Agnostic to specific OS or discovery stages. |
| `internal/discovery` | Specific capability implementations. | Flat `stage_*.go` files only. No subdirectories. |
| `internal/platform` | OS, Hardware, Network abstractions. | Use Go build tags (`provider_linux.go`). |
| `internal/services` | Infrastructure dependencies. | Completely isolated behind service interfaces. |
| `internal/presentation` | Output formatting (JSON, Text). | Depends on core artifacts, not implementations. |

## Architectural Invariants

These invariants must be preserved at all times:
- `internal/core` must never depend on any other `internal/*` package (except standard utilities like `domain`).
- The `discovery` package must remain flat. Adding a new stage means adding `stage_new.go`, not a new subdirectory.
- The `platform` package must use build tags for OS-specific logic. Do not create `linux/` or `windows/` directories.
- `internal/services` must never leak its third-party dependencies (e.g., PocketBase SDK types) into `core` or `discovery`.
- Global state (e.g., `runtime`) must not be imported by inner packages to bypass the `bootstrap` injection flow.

## When Architecture May Be Changed

Architecture may only be modified when:
- Introducing an entirely new top-level domain (e.g., introducing the `planner` or `executor` for v0.5.0/v0.6.0).
- An existing layer boundary physically prevents the implementation of a critical feature, and the change has been formally proposed and accepted.

## When Architecture Must NOT Be Changed

Architecture must **NOT** be changed for:
- Aesthetic preferences (e.g., renaming `core` to `engine` because it sounds better).
- Grouping files arbitrarily to "reduce file count" in a directory when it breaks cohesion.
- Bypassing dependency injection for convenience (e.g., introducing a global variable to avoid passing a context).
- Moving packages simply to satisfy a subjective view of clean architecture that contradicts the established `FREEZE.md`.
