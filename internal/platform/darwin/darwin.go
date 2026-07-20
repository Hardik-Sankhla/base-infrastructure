package darwin

import (
	"context"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform"
)

// Platform implements platform.Platform for macOS.
type Platform struct {
	osProvider platform.OSProvider
}

func NewPlatform() *Platform {
	return &Platform{
		osProvider: &OSProviderStub{},
	}
}

func (p *Platform) ID() string              { return "darwin" }
func (p *Platform) Name() string            { return "macOS" }
func (p *Platform) OS() platform.OSProvider { return p.osProvider }

type OSProviderStub struct{}

func (s *OSProviderStub) GetOSInfo(ctx context.Context) (models.OSInfo, error) {
	return models.OSInfo{OperatingSystem: "darwin"}, nil
}
