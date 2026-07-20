package windows

import "github.com/base-infrastructure/platform/internal/platform"

// Platform implements platform.Platform for Windows.
type Platform struct {
	osProvider platform.OSProvider
}

// NewPlatform creates a new Windows platform instance.
func NewPlatform() *Platform {
	return &Platform{
		osProvider: NewOSProvider(),
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
