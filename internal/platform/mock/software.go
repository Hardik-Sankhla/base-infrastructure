package mock

import (
	"github.com/base-infrastructure/platform/internal/runtime"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

type SoftwareProvider struct {
	Info models.SoftwareInfo
	Err  error
}

func (p *SoftwareProvider) GetSoftwareInfo(ctx runtime.Context) (models.SoftwareInfo, error) {
	return p.Info, p.Err
}
