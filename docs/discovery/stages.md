# Discovery Stages

A `Stage` is an isolated module responsible for interrogating a very specific slice of the environment.

## The Stage Interface
Every stage must adhere to `discovery.Stage`:

- `Name() string`: Unique identifier for the stage.
- `Version() string`: Semantic version of the stage logic.
- `Description() string`: Human-readable summary.
- `Priority() int`: Defines relative run ordering if dependencies aren't explicit.
- `DependsOn() []string`: A list of stage names that MUST successfully complete before this stage can begin.
- `Timeout() time.Duration`: Maximum allowed execution time before forced termination.
- `Initialize(context.Context) error`: Pre-execution setup hook.
- `Run(context.Context) (DiscoveryArtifact, error)`: The core execution logic.
- `Validate(DiscoveryArtifact) error`: Validates the integrity of the output.
- `Cleanup(context.Context) error`: Post-execution teardown hook.

## Built-in Stages
Located in `internal/discovery/builtin/stages.go`:

1. **OS Stage** (`internal/discovery/os`): Discovers the Host OS Name, Version, Kernel, and Architecture.
2. **Hardware Stage** (`internal/discovery/hardware`): Identifies CPU layout, Memory constraints, and virtualized container environments (e.g. Docker, LXC, WSL2).
3. **Network Stage** (`internal/discovery/network`): Pulls active interfaces, IPs, and egress capabilities.
4. **Filesystem Stage** (`internal/discovery/filesystem`): Mounts, ephemeral volumes, and disk capacity limits.
5. **Environment Stage** (`internal/discovery/environment`): Active environment variables, shell properties, and user context.
6. **Software Stage** (`internal/discovery/software`): Locates critical binaries required for capabilities (e.g., Python, Node, Git, Docker).

