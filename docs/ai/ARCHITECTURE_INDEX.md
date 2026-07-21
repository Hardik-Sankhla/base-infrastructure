# ARCHITECTURE_INDEX.md
# Architecture Index — Interface → Implementation → Test

> Every entry is source-verified with file path and line number.
> **Last scan**: 2026-07-21

---

## Engine Contracts

### `contracts.DiscoveryEngine`
- **Defined**: `internal/domain/contracts/engines.go:9-12`
- **Implemented by**: `discovery.DefaultDiscoveryEngine` (`internal/discovery/engine.go:12`)
- **Constructor**: `discovery.NewDiscoveryEngine(registry, cfg)` (`engine.go:19`)
- **Entry point**: `engine.Run(pctx)` (`engine.go:29`)
- **Tests**: No dedicated engine test — tested via integration in `cmd/`

### `contracts.PlannerEngine`
- **Defined**: `internal/domain/contracts/engines.go:15-17`
- **Implemented by**: NONE — no struct found
- **Status**: 🔲 Planned

### `contracts.ExecutorEngine`
- **Defined**: `internal/domain/contracts/engines.go:20-22`
- **Implemented by**: NONE — no struct found
- **Status**: 🔲 Planned

### `contracts.ValidatorEngine`
- **Defined**: `internal/domain/contracts/engines.go:25-27`
- **Implemented by**: NONE — no struct found
- **Status**: 🔲 Planned

---

## Discovery Pipeline

### `discovery.Stage` Interface
- **Defined**: `internal/discovery/stage.go`
- **Implemented by** (all in `internal/discovery/`):
  - `os.Stage` → `discovery/os/stage.go`
  - `hardware.Stage` → `discovery/hardware/stage.go`
  - `network.Stage` → `discovery/network/stage.go`
  - `filesystem.Stage` → `discovery/filesystem/stage.go`
  - `environment.Stage` → `discovery/environment/stage.go`
  - `software.Stage` → `discovery/software/stage.go`
- **Registered**: `internal/discovery/builtin/stages.go`

### `discovery.Pipeline`
- **Defined**: `internal/discovery/pipeline.go:26`
- **Sorting**: Priority-based insertion sort (`pipeline.go:253-270`)
- **Cycle detection**: Calls `Validator.Validate()` before run (`pipeline.go:74`)
- **Stage lifecycle per run**: `Initialize` → `Run` → `Validate` → `Cleanup` (`pipeline.go:126-147`)
- **FailFast behavior**: If `config.FailFast == true`, pipeline aborts on first stage error (`pipeline.go:103`)

### `discovery.Validator`
- **Defined**: `internal/discovery/validator.go:8`
- **Algorithm**: Depth-First Search (DFS) via recursive `hasCycle` closure (`validator.go:43`)
- **Checks**: duplicate names (`validator.go:22`), missing deps (`validator.go:30`), circular deps (`validator.go:43`)
- **Tests**: `internal/discovery/validator_test.go` — 4 test cases covering all scenarios

---

## Platform Abstraction

### `platform.Platform` Interface
- **Defined**: `internal/platform/platform.go:54-69`
- **Implemented by**:
  - `linux.Platform` → `internal/platform/linux/linux.go`
  - `windows.Platform` → `internal/platform/windows/windows.go`
  - `darwin.Platform` → `internal/platform/darwin/darwin.go`
  - `android.Platform` → `internal/platform/android/android.go`
  - `bsd.Platform` → `internal/platform/bsd/bsd.go`
  - `mock.Platform` → `internal/platform/mock/mock.go` (testing only)

### `platform.Detector` Interface
- **Defined**: `internal/platform/platform.go:72-75`
- **Implemented by**: `detector.Detector` → `internal/platform/detector/detector.go`

---

## Capability System

### `capabilities.Builder`
- **Defined**: `internal/capabilities/builder.go:10`
- **Input**: `*models.DiscoveryManifest`
- **Output**: `[]models.Capability`
- **Evaluated artifacts**: `"network"` (`builder.go:33`), `"software"` (`builder.go:53`)
- **NOT yet evaluated**: `"os"`, `"hardware"`, `"filesystem"`, `"environment"`
- **Tests**: `internal/capabilities/builder_test.go`

### `models.Capability`
- **Defined**: `internal/domain/models/capability.go:14`
- **Fields**: `ID`, `Provider`, `Version`, `State` (`StateAvailable|StateMissing|StateBroken`), `Confidence` (0-100), `Metadata`

---

## Runtime System

### `context.PlatformContext`
- **Defined**: `internal/runtime/context/context.go:17`
- **Constructor**: `NewPlatformContext(cfg, db)` (`context.go:31`)
- **Fields**: `Logger`, `Config`, `DB`, `EventBus`, `TaskEngine`, `FS`, `Downloader`, `Registry`, `goCtx`

### `events.Bus`
- **Defined**: `internal/runtime/events/bus.go`
- **Used for**: `DiscoveryStarted`, `DiscoveryFinished`, `DiscoveryStageStarted`, `DiscoveryStageCompleted`, `DiscoveryStageFailed`

### `plugin.Manifest`
- **Defined**: `internal/runtime/plugin/manifest.go:10`
- **Required fields**: `schema_version`, `name`, `version`
- **Loader**: `LoadManifest(path string) (*Manifest, error)` (`manifest.go:32`)

---

## CLI Commands

### `platform bootstrap`
- **Source**: `cmd/platform/cmd/bootstrap.go:16`
- **Flow**: RegisterStages → NewDiscoveryEngine → Run → NewBuilder → Build → JSON output

### `platform sdk validate [path]`
- **Source**: `cmd/platform/cmd/sdk.go:23`
- **Argument**: positional `[path]` via `cobra.ExactArgs(1)` — NOT `--path` flag
- **Note**: Currently only prints path, does NOT call `LoadManifest()` — validation is a stub

### `platform sdk create-plugin [name]`
- **Source**: `cmd/platform/cmd/sdk.go:14`
- **Status**: Stub — prints name only

### `platform sdk test [path]`
- **Source**: `cmd/platform/cmd/sdk.go:30`
- **Status**: Stub — prints path only
