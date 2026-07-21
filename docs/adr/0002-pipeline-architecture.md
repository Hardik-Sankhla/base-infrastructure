# ADR 0002: Pipeline Architecture for Discovery

## Status
Accepted

## Context
When bootstrapping a machine, we must discover its OS, hardware, network interfaces, installed software, and user permissions. Querying this sequentially takes an extremely long time. Querying it randomly results in errors (e.g. attempting to search for Windows registry keys before verifying the OS is actually Windows).

## Decision
We implemented a Priority-sorted Pipeline for the Discovery Engine.
Stages declare their explicit dependencies (`DependsOn() []string`). 
The engine uses Depth-First Search (DFS) to validate there are no circular loops, and then sorts the stages based on Priority. Stages execute sequentially based on priority ordering.

## Consequences

### Positive
- **Safety**: Impossible to run dependent stages out of order.
- **Speed**: Massive performance gain due to concurrent stage execution.
- **Resilience**: If the `Network` stage fails, stages depending on it are cleanly aborted, but unrelated stages (like `Hardware`) continue executing.

### Negative
- **Complexity**: Writing a new stage requires understanding how to properly declare dependencies to avoid deadlocks.

