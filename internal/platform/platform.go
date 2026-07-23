package platform

import (
	"github.com/base-infrastructure/platform/internal/runtime"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// HardwareProvider abstracts hardware resource discovery.
type HardwareProvider interface {
	GetCPU(ctx runtime.Context) (models.CPU, error)
	GetRAM(ctx runtime.Context) (models.RAM, error)
	GetStorage(ctx runtime.Context) ([]models.Disk, error)
	GetGPUs(ctx runtime.Context) ([]models.GPU, error)
	GetBattery(ctx runtime.Context) (models.Battery, error)
	GetThermal(ctx runtime.Context) ([]models.ThermalSensor, error)
}

// OSProvider abstracts the retrieval of operating system information.
type OSProvider interface {
	GetOSInfo(ctx runtime.Context) (models.OSInfo, error)
}

// FilesystemProvider abstracts filesystem context.
type FilesystemProvider interface {
	GetFilesystemInfo(ctx runtime.Context) (models.FilesystemInfo, error)
}

// NetworkProvider abstracts the retrieval of network configuration and state.
type NetworkProvider interface {
	GetInterfaces(ctx runtime.Context) ([]models.NetworkInterface, error)
	GetDNS(ctx runtime.Context) (models.DNSConfig, error)
	GetProxy(ctx runtime.Context) (models.ProxyConfig, error)
}

// EnvironmentProvider abstracts the retrieval of execution environment context.
type EnvironmentProvider interface {
	GetEnvironmentInfo(ctx runtime.Context) (models.EnvironmentInfo, error)
}

// SoftwareProvider abstracts the retrieval of installed software and runtimes.
type SoftwareProvider interface {
	GetSoftwareInfo(ctx runtime.Context) (models.SoftwareInfo, error)
}

// Other providers for the future
type (
	ProcessProvider  interface{}
	SecurityProvider interface{}
	UserProvider     interface{}
	ServiceProvider  interface{}
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
