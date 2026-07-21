# TESTING_RULES.md
# Testing Standards & Rules

> Derived from source analysis of test files and CI configuration.
> **Last verified**: 2026-07-21

---

## Mandatory Testing Rules

### Rule T1: Every Stage Requires a `_test.go` File
All 6 built-in discovery stages have companion test files (verified):
- `internal/discovery/environment/stage_test.go`
- `internal/discovery/filesystem/stage_test.go`
- `internal/discovery/hardware/stage_test.go`
- `internal/discovery/network/stage_test.go`
- `internal/discovery/os/stage_test.go`
- `internal/discovery/software/stage_test.go`

Any new stage added to `internal/discovery/` MUST have a `stage_test.go`.

### Rule T2: Discovery Stage Tests MUST Use MockPlatform
Physical OS discovery commands (e.g., `wmic`, `sysctl`, reading `/proc/cpuinfo`) are not available in GitHub Actions CI environments on all targets.

All stage tests MUST use `internal/platform/mock/` to provide fake platform data:
```go
import "github.com/base-infrastructure/platform/internal/platform/mock"

plat := mock.NewMockPlatform()
dctx := discovery.NewContext(logger, bus, cfg, db, plat)
```

### Rule T3: Race Detector is Mandatory in CI
The CI workflow runs `go test -v -race ./...` (source: `.github/workflows/ci.yml:62`).

Locally, always run:
```bash
go test -race ./...
```

The pipeline uses goroutines for concurrent stage execution. Race conditions are a real risk.

### Rule T4: Validator Tests Cover All DAG Failure Modes
`internal/discovery/validator_test.go` has 4 test functions (verified):
- `TestValidator_ValidGraph`
- `TestValidator_DuplicateStage`
- `TestValidator_MissingDependency`
- `TestValidator_CircularDependency`

Any new validator logic MUST add a corresponding test case.

### Rule T5: Test Fakes Live in `internal/testing/fakes/`
Do not create ad-hoc mock structs inline in test files. Use or extend the shared fakes:
- `internal/testing/fakes/downloader.go`
- `internal/testing/fakes/fs.go`

---

## Current Coverage Gaps

| Package | Coverage | Action Required |
|---|---|---|
| `cmd/platform` | 0.0% | Add integration test that runs `bootstrap` command |
| `internal/config` | 0.0% | Add unit tests for config loading |
| `internal/capabilities/builder.go` | UNVERIFIED | Tests exist (`builder_test.go`) but coverage % unknown |
| `internal/discovery` (pipeline/engine) | Low | Add pipeline-level integration tests |
| `internal/runtime/` | UNVERIFIED | Coverage unknown |

---

## Running Tests Locally

```bash
# All tests
go test ./...

# With race detector (mandatory before any PR)
go test -race ./...

# With coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Specific package
go test ./internal/discovery/...
```
