# Glossary & Terminology

### Artifact (DiscoveryArtifact)
A single piece of data returned by a Discovery Stage. Must implement `ArtifactType() string` to allow dynamic casting.

### Capability
A discrete, standardized requirement that has been fulfilled by the host environment (e.g., `runtime.docker`).

### DAG (Directed Acyclic Graph)
A computational graph where stages run in a specific dependency order without circular loops. The Discovery Engine uses a DAG to safely organize execution (e.g., Network discovery cannot start until OS discovery finishes).

### Discovery Engine
The overarching runner that manages the entire lifecycle of interrogating the environment.

### Hook
A middleware event registered in the Discovery Engine (e.g., `PreRun`, `PostRun`). Allows injection of loggers or custom validations.

### Manifest (DiscoveryManifest)
The final aggregated output of the entire Discovery Pipeline.

### Platform
The abstracted operating system layer. The Go runtime interfaces with the `Platform` which translates commands down to `Windows`, `Linux`, `Darwin`, etc.

### Plugin
An external script (usually a bash or PowerShell script) executed via JSON-RPC that can install software to fulfill missing Capabilities.

### Stage
An isolated interrogation module (e.g., `SoftwareStage`). It must implement `Initialize`, `Run`, `Validate`, and `Cleanup`.

