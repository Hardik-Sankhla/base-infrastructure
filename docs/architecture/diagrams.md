# System Diagrams

## Repository Architecture
```mermaid
graph TD
    CLI[cmd/platform] --> Builder[internal/capabilities]
    CLI --> Pipeline[internal/discovery]
    
    Pipeline --> StageOS[OS Stage]
    Pipeline --> StageHW[Hardware Stage]
    Pipeline --> StageNet[Network Stage]
    Pipeline --> StageEnv[Environment Stage]
    Pipeline --> StageFS[Filesystem Stage]
    Pipeline --> StageSoft[Software Stage]
    
    StageOS --> PlatformAbstraction[internal/platform]
    StageHW --> PlatformAbstraction
    StageNet --> PlatformAbstraction
    StageEnv --> PlatformAbstraction
    StageFS --> PlatformAbstraction
    StageSoft --> PlatformAbstraction

    PlatformAbstraction --> HostOS((Host System))
```

## Artifact Flow
```mermaid
sequenceDiagram
    participant CLI
    participant DiscoveryEngine
    participant Stage
    participant Builder
    
    CLI->>DiscoveryEngine: Run(Context)
    DiscoveryEngine->>Stage: Execute Stage
    Stage->>DiscoveryEngine: Return DiscoveryArtifact
    DiscoveryEngine->>DiscoveryEngine: Store in Manifest Map
    DiscoveryEngine-->>CLI: Return DiscoveryManifest
    CLI->>Builder: NewBuilder(Manifest)
    Builder->>CLI: Return []Capability
```

## Discovery Pipeline Lifecycle
```mermaid
stateDiagram-v2
    [*] --> Initialization
    Initialization --> DependencyResolution
    DependencyResolution --> ParallelExecution
    ParallelExecution --> StageExecution
    
    state StageExecution {
        [*] --> Setup
        Setup --> Run
        Run --> Validate
        Validate --> Cleanup
        Cleanup --> [*]
    }
    
    StageExecution --> MergeArtifacts
    MergeArtifacts --> [*]
```

