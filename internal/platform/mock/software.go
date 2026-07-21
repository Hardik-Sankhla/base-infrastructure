package mock

import (
	"context"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

type SoftwareProvider struct {
	Info models.SoftwareInfo
	Err  error
}

func (p *SoftwareProvider) GetSoftwareInfo(ctx context.Context) (models.SoftwareInfo, error) {
	return p.Info, p.Err
}
