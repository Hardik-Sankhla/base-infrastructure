# Contracts & Interfaces

The interfaces that define the architectural boundaries of the system.

## Discovery Stage (`discovery.Stage`)
Defined in `internal/discovery/stage.go`.
```go
type Stage interface {
	Name() string
	Version() string
	Description() string
	Priority() int
	DependsOn() []string
	Timeout() time.Duration
	Initialize(context.Context) error
	Run(context.Context) (DiscoveryArtifact, error)
	Validate(DiscoveryArtifact) error
	Cleanup(context.Context) error
}
```

## Platform (`platform.Platform`)
Defined in `internal/platform/platform.go`.
```go
type Platform interface {
	Name() string
	Version() string
	OS() OSProvider
	Hardware() HardwareProvider
	Network() NetworkProvider
	Environment() EnvironmentProvider
	Filesystem() FilesystemProvider
	Software() SoftwareProvider
	Close() error
}
```

## Engines (`contracts.go`)
Located at `internal/domain/contracts/engines.go`. Provides abstract engine lifecycle hooks.

