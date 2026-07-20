package bsd

import (
	"context"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform"
)

type Platform struct {
	osProvider platform.OSProvider
}

func NewPlatform() *Platform {
	return &Platform{osProvider: &OSProviderStub{}}
}

func (p *Platform) ID() string              { return "bsd" }
func (p *Platform) Name() string            { return "BSD" }
func (p *Platform) OS() platform.OSProvider { return p.osProvider }

type OSProviderStub struct{}

func (s *OSProviderStub) GetOSInfo(ctx context.Context) (models.OSInfo, error) {
	return models.OSInfo{OperatingSystem: "bsd"}, nil
}
