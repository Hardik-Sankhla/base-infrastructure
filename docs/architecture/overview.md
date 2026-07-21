# Architecture Overview

The Base Infrastructure Platform is built around a decoupled architecture. It strictly separates the act of *interrogating* an environment (Discovery Engine) from the act of *mutating* an environment (Execution Engine).

## Core Philosophy
We treat an operating system environment not by its name, but by its verifiable **Capabilities**. For example, rather than executing scripts that assume Ubuntu because `uname` returns Linux, the system proves `apt` exists via the `SoftwareProvider`, turning it into a capability.

## Subsystems (Implemented)

### 1. Platform Abstraction Layer
Located at `internal/platform/`. This acts as the lowest-level boundary interface between the Go runtime and the Host OS.
- **Purpose**: Normalize disparate operating system behaviors.
- **Responsibilities**: Execute precise, environment-specific commands (like `wmic` on Windows or `sysctl` on macOS) to fetch hardware, os, and network data.
- **Public Interfaces**: `Provider` interface grouped by subsystems (e.g. `HardwareProvider`, `OSProvider`).

### 2. Discovery Engine
Located at `internal/discovery/`.
- **Purpose**: A concurrent Directed Acyclic Graph (DAG) executor that runs interrogation plugins/stages safely.
- **Responsibilities**: Sorting stages by dependencies, enforcing timeouts, caching artifact outputs, and handling context propagation.
- **Design Rationale**: A pipeline architecture allows discovery stages to safely depend on one another without blocking the entire execution thread. For instance, the `Network` stage relies on the `OS` stage completing first.

### 3. Capability Builder
Located at `internal/capabilities/`.
- **Purpose**: Transforms heterogeneous `DiscoveryManifest` artifacts into homogenous Platform `Capabilities`.
- **Responsibilities**: Reading unstructured map data provided by the discovery pipeline and asserting types into standardized `models.Capability` items (e.g., ID: `network.connectivity`).

### 4. Domain Models
Located at `internal/domain/models/`.
- **Purpose**: Expose the structural bounds of all data flowing through the runtime.
- **Design Rationale**: Isolating models prevents circular dependencies and ensures strict contract enforcement between the engine and the platform.

## Subsystems (Planned)

### 1. Execution Engine
- **Purpose**: Receive the desired state payload, diff it against the `Capabilities` list, and execute task hooks.
- **Lifecycle**: Calculates state drift -> Builds execution graph -> Triggers plugins.

### 2. Plugin JSON-RPC Host
- **Purpose**: Safely invoke arbitrary scripts (Bash, PowerShell, Python) with STDIN payloads.
- **Current Status**: Alpha manifest specifications located in `pkg/sdk/` and `plugins/`.

