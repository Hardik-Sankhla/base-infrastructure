# ADR 0003: Platform Abstraction Layer

## Status
Accepted

## Context
If the core pipeline directly queries `wmic` for hardware, the code becomes littered with `if runtime.GOOS == "windows"` blocks. As we expand to Android, Darwin, BSD, and esoteric embedded systems, the cyclomatic complexity of the core pipeline would become unmaintainable.

## Decision
We introduced a strict Platform Abstraction Layer (`internal/platform`).
At boot, a `Detector` determines the host OS and returns an object satisfying the `Platform` interface. 
This interface provides grouped "Providers" (e.g. `OSProvider`, `NetworkProvider`) which execute native commands safely.

## Consequences

### Positive
- **Decoupling**: The Discovery Engine has zero knowledge of how an IP address is fetched. It just calls `Platform.Network().DiscoverInterfaces()`.
- **Testability**: We created a `MockPlatform` which satisfies the entire interface. The Discovery Engine can be 100% unit tested without touching a physical kernel.

