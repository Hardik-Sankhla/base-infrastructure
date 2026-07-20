package linux

import "github.com/base-infrastructure/platform/internal/platform"

// Platform implements platform.Platform for Linux.
type Platform struct {
	osProvider platform.OSProvider
}

// NewPlatform creates a new Linux platform instance.
func NewPlatform() *Platform {
	return &Platform{
		osProvider: NewOSProvider(),
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
