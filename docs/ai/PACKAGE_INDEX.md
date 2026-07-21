# PACKAGE_INDEX.md
# Package Reference Index

> Source-verified. Every package is confirmed by directory listing.
> **Last scan**: 2026-07-21

---

## Binary Entrypoint

### `cmd/platform/`
The CLI binary entrypoint. Uses [Cobra](https://github.com/spf13/cobra).

| File | Purpose |
|---|---|
| `cmd/platform/main.go` | `main()` — calls `cmd.Execute()` |
| `cmd/platform/cmd/root.go` | Root Cobra command, global flags, config loading |
| `cmd/platform/cmd/bootstrap.go` | `platform bootstrap` — runs discovery + capability build |
| `cmd/platform/cmd/sdk.go` | `platform sdk` — validate, create-plugin, test subcommands |

---

## Internal Packages (`internal/`)

### `internal/capabilities/`
Translates `DiscoveryManifest` → `[]models.Capability`.
- `builder.go` — `Builder` struct with `Build()`, `evaluateNetwork()`, `evaluateSoftware()`
- `builder_test.go` — Unit tests for capability generation

### `internal/config/`
Configuration loading via [Viper](https://github.com/spf13/viper).
- `config.go` — `Config` struct, default values, YAML loading

### `internal/discovery/`
The core pipeline execution engine.
- `artifact.go` — `DiscoveryArtifact` interface
- `cache.go` — Thread-safe artifact cache (`sync.RWMutex`)
- `context.go` — Discovery-specific context (wraps `PlatformContext`)
- `doc.go` — Package documentation
- `engine.go` — `DefaultDiscoveryEngine` — orchestrates detection + pipeline
- `hooks.go` — `Hook` interface for pipeline lifecycle events
- `pipeline.go` — `Pipeline` — stage sorting, execution, fail-fast
- `registry.go` — `Registry` — holds registered stages
- `result.go` — `Result`, `StageResult`, `ResultBuilder`
- `stage.go` — `Stage` interface (10 methods)
- `validator.go` — DFS cycle + duplicate + missing-dep validator
- `validator_test.go` — 4 DAG tests

**Sub-packages** (each implements `Stage`):
- `discovery/builtin/stages.go` — registers all core stages
- `discovery/environment/stage.go` + `stage_test.go`
- `discovery/filesystem/stage.go` + `stage_test.go`
- `discovery/hardware/stage.go` + `stage_test.go`
- `discovery/network/stage.go` + `stage_test.go`
- `discovery/os/stage.go` + `stage_test.go`
- `discovery/software/stage.go` + `stage_test.go`

### `internal/domain/`
Pure domain — no external dependencies allowed here.

**`internal/domain/contracts/`**
- `engines.go` — `DiscoveryEngine`, `PlannerEngine`, `ExecutorEngine`, `ValidatorEngine`

**`internal/domain/models/`**
- `capability.go` — `Capability`, `CapabilityState`
- `compatibility.go` — OS/arch compatibility structs
- `discovery.go` — `DiscoveryManifest`, `StageExecutionResult`
- `environment.go` — `EnvironmentInfo`
- `network.go` — `NetworkInterface`, `DNSConfig`, `ProxyConfig`
- `plan.go` — `ExecutionPlan`, `PlanTask`
- `policy.go` — `Policy` struct
- `resource.go` — `Resource` struct
- `result.go` — `Result` struct
- `security.go` — `SecurityInfo`
- `software.go` — `SoftwareInfo`, `RuntimeEnvironment`, `Tool`

### `internal/executor/`
**Status**: 🔲 Package exists — content UNVERIFIED (not fully read)

### `internal/logger/`
- `logger.go` — Structured logging setup (`log/slog`)

### `internal/planner/`
**Status**: 🔲 Package exists — content UNVERIFIED

### `internal/platform/`
OS-specific implementations. See `ARCHITECTURE_INDEX.md`.

### `internal/runtime/`
Runtime infrastructure services.
- `runtime/context/context.go` — `PlatformContext` (DI root)
- `runtime/events/bus.go` — Pub/sub event bus
- `runtime/fs/manager.go` — Filesystem abstraction
- `runtime/http/downloader.go` — HTTP download utility
- `runtime/plugin/manifest.go` — Plugin manifest loader
- `runtime/plugin/runtime.go` — Plugin runtime (UNVERIFIED)
- `runtime/tasks/engine.go` — Background task engine

### `internal/state/`
- `db.go` — SQLite state persistence via `modernc.org/sqlite`

### `internal/testing/`
Test helpers and fakes.
- `testing/fakes/downloader.go` — Fake HTTP downloader
- `testing/fakes/fs.go` — Fake filesystem

---

## Public Packages (`pkg/`)

### `pkg/sdk/`
**Status**: UNVERIFIED — directory exists, content not yet read

---

## External Dependencies (Key)

| Dependency | Version | Purpose |
|---|---|---|
| `github.com/spf13/cobra` | v1.8.0 | CLI framework |
| `github.com/spf13/viper` | v1.18.2 | Config loading |
| `github.com/shirou/gopsutil/v3` | v3.24.5 | Hardware metrics |
| `modernc.org/sqlite` | v1.29.5 | Embedded state database |
| `gopkg.in/yaml.v3` | v3.0.1 | YAML parsing |
