# REPOSITORY_STATE.md
# Current Implementation State — Source-Verified

> All claims in this document are backed by source file + line number.
> **UNVERIFIED** means a claim could not be confirmed from source.
> **Last full scan**: 2026-07-21

---

## Engine Contracts (`internal/domain/contracts/engines.go`)

| Contract | Status | Source Evidence |
|---|---|---|
| `DiscoveryEngine.Run()` | ✅ Implemented | `engines.go:9-12` → `internal/discovery/engine.go:29` |
| `PlannerEngine.Plan()` | 🔲 Interface only | `engines.go:15-17` — no implementing struct found |
| `ExecutorEngine.Execute()` | 🔲 Interface only | `engines.go:20-22` — no implementing struct found |
| `ValidatorEngine.Validate()` | 🔲 Interface only | `engines.go:25-27` — no implementing struct found |

---

## Platform Providers (`internal/platform/platform.go`)

| Provider | Methods | OS Implementations | Status |
|---|---|---|---|
| `HardwareProvider` | `GetCPU`, `GetRAM`, `GetStorage`, `GetGPUs`, `GetBattery`, `GetThermal` | linux, windows, darwin, android, bsd (via `providers/hardware`) | ✅ Implemented |
| `OSProvider` | `GetOSInfo` | linux (`linux/os.go`), windows (`windows/os.go`) | ✅ Implemented |
| `FilesystemProvider` | `GetFilesystemInfo` | linux (`linux/fs.go`), windows (`windows/fs.go`) | ✅ Implemented |
| `NetworkProvider` | `GetInterfaces`, `GetDNS`, `GetProxy` | linux, windows, darwin, android, bsd | ✅ Implemented |
| `EnvironmentProvider` | `GetEnvironmentInfo` | linux, windows, darwin, android, bsd | ✅ Implemented |
| `SoftwareProvider` | `GetSoftwareInfo` | linux, windows, darwin, android, bsd | ✅ Implemented |
| `ProcessProvider` | _(none)_ | _(none)_ | 🔲 Empty stub (`platform.go:48`) |
| `SecurityProvider` | _(none)_ | _(none)_ | 🔲 Empty stub (`platform.go:49`) |
| `UserProvider` | _(none)_ | _(none)_ | 🔲 Empty stub (`platform.go:50`) |
| `ServiceProvider` | _(none)_ | _(none)_ | 🔲 Empty stub (`platform.go:51`) |

---

## Discovery Stages (`internal/discovery/`)

| Stage Package | Stage Name | Has Test | Source |
|---|---|---|---|
| `discovery/os` | `os` | ✅ `stage_test.go` | `internal/discovery/os/stage.go` |
| `discovery/hardware` | `hardware` | ✅ `stage_test.go` | `internal/discovery/hardware/stage.go` |
| `discovery/network` | `network` | ✅ `stage_test.go` | `internal/discovery/network/stage.go` |
| `discovery/filesystem` | `filesystem` | ✅ `stage_test.go` | `internal/discovery/filesystem/stage.go` |
| `discovery/environment` | `environment` | ✅ `stage_test.go` | `internal/discovery/environment/stage.go` |
| `discovery/software` | `software` | ✅ `stage_test.go` | `internal/discovery/software/stage.go` |

---

## Capability Builder (`internal/capabilities/builder.go`)

Currently produces capabilities from exactly **2** artifact types:

| Artifact Key | Capability IDs Generated | Logic | Source |
|---|---|---|---|
| `"network"` | `network.connectivity` | If any interface is `IsUp && len(IPv4) > 0` | `builder.go:36-44` |
| `"software"` | `runtime.<name>` for each runtime | Always | `builder.go:57` |
| `"software"` | `container.runtime` | Only if `runtime.name == "docker"` | `builder.go:64-72` |

**Not yet evaluated**: `os`, `hardware`, `filesystem`, `environment` artifacts.

---

## PlatformContext (`internal/runtime/context/context.go`)

Constructed via `NewPlatformContext(cfg, db)` at `context.go:31`.

| Field | Type | Purpose |
|---|---|---|
| `Logger` | `*slog.Logger` | Structured logging |
| `Config` | `*config.Config` | Loaded from YAML via Viper |
| `DB` | `*sql.DB` | SQLite via `modernc.org/sqlite` |
| `EventBus` | `events.Bus` | Pub/sub event system |
| `TaskEngine` | `tasks.Engine` | Background task runner |
| `FS` | `fs.Manager` | Filesystem abstraction |
| `Downloader` | `http.Downloader` | HTTP download utility |
| `Registry` | `plugin.Registry` | Loaded plugin manifests |

---

## Plugin System

| Component | Status | Source |
|---|---|---|
| `Manifest` struct | ✅ Implemented | `internal/runtime/plugin/manifest.go:10` |
| `LoadManifest(path)` | ✅ Implemented | `manifest.go:32` — validates `schema_version`, `name`, `version` |
| Plugin execution (subprocess STDIN/STDOUT) | 🔲 NOT implemented | UNVERIFIED — no executor found |
| `platform sdk validate` | ✅ CLI works | `cmd/platform/cmd/sdk.go` — positional arg, NOT `--path` flag |

### ⚠️ Known Schema Mismatch
The bundled plugins (`plugins/*/manifest.yaml`) are missing `schema_version`.
Running `platform sdk validate plugins/docker/manifest.yaml` will **return an error**.
See [`KNOWN_ISSUES.md`](./KNOWN_ISSUES.md#plugin-manifest-schema-mismatch).

---

## Test Coverage (from CI logs, 2026-07-21)

| Package | Coverage | Note |
|---|---|---|
| `internal/discovery` (validator) | ~13.4% | DAG validation tests pass |
| `cmd/platform` | 0.0% | No unit tests |
| `internal/config` | 0.0% | No unit tests |
| All stage `_test.go` files | Present | Verified by file listing |
