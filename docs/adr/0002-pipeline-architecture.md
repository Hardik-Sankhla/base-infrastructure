# ADR 0002: Pipeline Architecture for Discovery

## Status
Accepted

## Context
When bootstrapping a machine, we must discover its OS, hardware, network interfaces, installed software, and user permissions. Querying this sequentially takes an extremely long time. Querying it randomly results in errors (e.g. attempting to search for Windows registry keys before verifying the OS is actually Windows).

## Decision
We implemented a Directed Acyclic Graph (DAG) Pipeline for the Discovery Engine.
Stages declare their explicit dependencies (`DependsOn() []string`). 
The engine uses Kahn's Topological Sort to group stages into tiers, running stages in the same tier concurrently, and blocking subsequent tiers until all prerequisites succeed.

## Consequences

### Positive
- **Safety**: Impossible to run dependent stages out of order.
- **Speed**: Massive performance gain due to concurrent stage execution.
- **Resilience**: If the `Network` stage fails, stages depending on it are cleanly aborted, but unrelated stages (like `Hardware`) continue executing.

### Negative
- **Complexity**: Writing a new stage requires understanding how to properly declare dependencies to avoid deadlocks.

