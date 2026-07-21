# Domain Models

The core structs used by the platform are located in `internal/domain/models/`.

## Capability (`models.Capability`)
A standardized capability that has been proven to exist in the environment.
- `ID string`: A dot-notation tag (e.g., `network.connectivity`, `runtime.docker`).
- `Provider string`: The specific engine fulfilling this capability.
- `Version string`: The semantic version.

## DiscoveryArtifact
An interface that all specific discovery models must implement to be legally passed through the engine pipeline.
- `ArtifactType() string`: Returns the unique string mapping for the struct (e.g. `"os"`, `"software"`).

## OSInfo (`models.OSInfo`)
Implements `DiscoveryArtifact`.
- `Family string`: Generic family (e.g. `linux`, `windows`).
- `Name string`: Specific distro (e.g. `ubuntu`).
- `Version string`: Distro version.
- `Architecture string`: CPU Arch (`amd64`, `arm64`).

## SoftwareInfo (`models.SoftwareInfo`)
Implements `DiscoveryArtifact`.
- `Runtimes []RuntimeEnvironment`: Contains dynamic properties (`Name`, `Version`, `Path`) for installed runtimes.
- `Tools []Tool`: CLI tools available in PATH.

## DiscoveryManifest (`models.DiscoveryManifest`)
The root object that aggregates the results of all individual `DiscoveryArtifact` structs.
- `Environment string`: Root environment token.
- `Artifacts map[string]any`: The dynamic map holding the artifact structs.

