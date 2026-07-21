# ADR 0004: Immutable Discovery Cache

## Status
Accepted

## Context
Multiple components might need the exact same piece of discovery information at different times during execution. Without caching, we would unnecessarily query the host OS multiple times (e.g., executing `wmic cpu get` multiple times), slowing down the bootstrap process and risking race conditions.

## Decision
We introduced a concurrent-safe `DiscoveryCache` module (`internal/discovery/cache.go`).
Once a stage validates its generated `DiscoveryArtifact`, the engine stores a deep copy of that artifact in the Cache using a `sync.RWMutex`.

## Consequences

### Positive
- **Performance**: O(1) retrieval for subsequent data checks. No redundant OS commands executed.
- **Safety**: The cache hands out copies or interfaces, ensuring stages cannot accidentally mutate global state while validating.

