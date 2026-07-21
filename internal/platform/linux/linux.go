package linux

import (
	"github.com/base-infrastructure/platform/internal/platform"
	"github.com/base-infrastructure/platform/internal/platform/providers/hardware"
)

// Platform implements platform.Platform for Linux.
type Platform struct {
	platform.BasePlatform
	osProvider       platform.OSProvider
	hardwareProvider platform.HardwareProvider
	fsProvider       platform.FilesystemProvider
}

// NewPlatform creates a new Linux platform instance.
func NewPlatform() *Platform {
	return &Platform{
		osProvider:       NewOSProvider(),
		hardwareProvider: hardware.NewDefaultProvider(),
		fsProvider:       NewFilesystemProvider(),
	}
}

func (p *Platform) ID() string {
	return "linux"
}

func (p *Platform) Name() string {
	return "Linux"
}

func (p *Platform) OS() platform.OSProvider {
	return p.osProvider
}

func (p *Platform) Hardware() platform.HardwareProvider {
	return p.hardwareProvider
}

func (p *Platform) Filesystem() platform.FilesystemProvider {
	return p.fsProvider
}

func (p *Platform) Network() platform.NetworkProvider {
	return NewNetworkProvider()
}

func (p *Platform) Environment() platform.EnvironmentProvider {
	return NewEnvironmentProvider()
}

func (p *Platform) Software() platform.SoftwareProvider {
	return NewSoftwareProvider()
}
