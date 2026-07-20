package platform

import (
	"context"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// HardwareProvider abstracts hardware resource discovery.
type HardwareProvider interface {
	GetCPU(ctx context.Context) (models.CPU, error)
	GetRAM(ctx context.Context) (models.RAM, error)
	GetStorage(ctx context.Context) ([]models.Disk, error)
	GetGPUs(ctx context.Context) ([]models.GPU, error)
	GetBattery(ctx context.Context) (models.Battery, error)
	GetThermal(ctx context.Context) ([]models.ThermalSensor, error)
}

// OSProvider abstracts the retrieval of operating system information.
type OSProvider interface {
	GetOSInfo(ctx context.Context) (models.OSInfo, error)
}

// FilesystemProvider abstracts filesystem context.
type FilesystemProvider interface {
	GetFilesystemInfo(ctx context.Context) (models.FilesystemInfo, error)
}

// Other providers for the future
type (
	NetworkProvider     interface{}
	EnvironmentProvider interface{}
	SoftwareProvider    interface{}
	ProcessProvider     interface{}
	SecurityProvider    interface{}
	UserProvider        interface{}
	ServiceProvider     interface{}
)

// Platform provides a cross-platform abstraction for discovery and execution.
type Platform interface {
	ID() string
	Name() string

	Hardware() HardwareProvider
	OS() OSProvider
	Filesystem() FilesystemProvider
	Network() NetworkProvider
	Environment() EnvironmentProvider
	Software() SoftwareProvider
	Process() ProcessProvider
	Security() SecurityProvider
	User() UserProvider
	Service() ServiceProvider
}

// Detector is responsible for identifying the current runtime environment
// and returning the appropriate Platform implementation.
type Detector interface {
	Detect() (Platform, error)
}
