package windows

import (
	"github.com/base-infrastructure/platform/internal/discovery/hardware"
	"github.com/base-infrastructure/platform/internal/platform"
)

// Platform implements platform.Platform for Windows.
type Platform struct {
	platform.BasePlatform
	osProvider       platform.OSProvider
	hardwareProvider platform.HardwareProvider
	fsProvider       platform.FilesystemProvider
}

// NewPlatform creates a new Windows platform instance.
func NewPlatform() *Platform {
	return &Platform{
		osProvider:       NewOSProvider(),
		hardwareProvider: hardware.NewDefaultProvider(),
		fsProvider:       NewFilesystemProvider(),
	}
}

func (p *Platform) ID() string {
	return "windows"
}

func (p *Platform) Name() string {
	return "Windows"
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
