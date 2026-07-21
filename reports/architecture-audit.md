# Certification: Architecture Audit

## 1. Clean Architecture & Boundaries
**Verified.** The core business logic (`internal/discovery/engine.go`) has zero dependencies on OS-specific packages.
- **Evidence**: `internal/discovery/engine.go:34` calls `detector.NewDetector().Detect()` which returns an abstract `Platform` interface. No `syscall` or `os/exec` is imported in `engine.go`.

## 2. Platform Abstraction
**Verified.** The OS-specific logic is strictly contained within `internal/platform/`.
- **Evidence**: `internal/platform/platform.go` defines the `Platform` interface. `internal/platform/mock/mock.go` implements this interface for testing.

## 3. Dependency Inversion
**Verified.** Modules depend on abstractions, not concretions.
- **Evidence**: `internal/discovery/stage.go` defines the `Stage` interface. The `Pipeline` in `internal/discovery/pipeline.go:88` accepts any slice of `[]Stage`.

## 4. Pipeline Execution Logic
**Correction.** Previous documentation claimed the pipeline used "Kahn's Topological Sort". This was an assumption and is **FALSE**.
- **Evidence (Cycle Detection)**: `internal/discovery/validator.go:43` uses a standard Depth-First Search (DFS) `hasCycle` recursive closure to detect dependency loops.
- **Evidence (Sorting)**: `internal/discovery/pipeline.go:253` uses a simple Priority-based bubble sort (`sortedStages()`), rather than a dynamic Kahn's algorithm graph traversal. 
- **Action**: Documentation must be updated to remove references to Kahn's Algorithm.

## 5. Caching
**Verified.** The `DiscoveryCache` utilizes `sync.RWMutex` to prevent concurrent map writes.
- **Evidence**: `internal/discovery/cache.go` (Verified in previous file inspection).
