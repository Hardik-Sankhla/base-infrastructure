# Platform Runtime

The Platform Runtime acts as the foundational layer binding all subsystems together.

## Context (`internal/runtime/context`)
The system passes a strongly-typed `discovery.Context` rather than a standard `context.Context` (though it inherits from it) throughout the execution lifecycle.

### Purpose
The custom Context tracks cancellation signals, timeout deadlines, and active Logger instances. It ensures that if a single Discovery Stage takes too long or hangs indefinitely (e.g. `wmic` failing on Windows), the entire stage can be forcefully preempted without crashing the core binary.

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

