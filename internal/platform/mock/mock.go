package mock

import (
	"context"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform"
)

type Platform struct {
	MockID   string
	MockName string
	MockOS   *OSProvider
}

func NewPlatform() *Platform {
	return &Platform{
		MockID:   "mock",
		MockName: "Mock OS",
		MockOS:   &OSProvider{},
	}
}

func (p *Platform) ID() string              { return p.MockID }
func (p *Platform) Name() string            { return p.MockName }
func (p *Platform) OS() platform.OSProvider { return p.MockOS }

type OSProvider struct {
	Info models.OSInfo
	Err  error
}

func (p *OSProvider) GetOSInfo(ctx context.Context) (models.OSInfo, error) {
	return p.Info, p.Err
}
