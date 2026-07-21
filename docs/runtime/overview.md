# Platform Runtime

The Platform Runtime acts as the foundational layer binding all subsystems together.

## Context (`internal/runtime/context`)
The system passes a unified strongly-typed `PlatformContext` rather than a standard `context.Context` (though it provides one via `GoContext()`) throughout the execution lifecycle.

### Purpose
The custom `PlatformContext` acts as the dependency injection (DI) root for the entire platform. It tracks cancellation signals, timeout deadlines, active Logger instances, configuration, and state persistence stores. The discovery engine also uses a specific `discovery.Context` facade that wraps `PlatformContext`.

## Events (`internal/runtime/events`)
The runtime implements an internal Event Bus for decoupled Pub/Sub communication.

### Purpose
To allow deep modules to broadcast their status without relying on rigid dependency injection. 
- A stage publishes `EventStageStarted`.
- The CLI logger subscribes to `EventStageStarted` and prints `[INFO] Running OS Stage...`.

## Tasks (`internal/runtime/tasks`)
(?? Planned for v0.2.0)

The Task Engine will be responsible for mapping desired state configurations to actionable execution commands.

## Hooks (`internal/discovery/hooks.go`)
Middleware functions inserted into the discovery lifecycle.
- `PreRun`: Executed before a stage invokes `Run()`.
- `PostRun`: Executed after a stage successfully completes `Run()`.
- `OnError`: Executed when a stage triggers an internal panic or returns an `error`.

