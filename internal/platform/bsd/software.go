package bsd

import (
	"context"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// SoftwareProvider implements platform.SoftwareProvider for BSD.
type SoftwareProvider struct{}

// NewSoftwareProvider creates a new BSD software provider.
func NewSoftwareProvider() *SoftwareProvider {
	return &SoftwareProvider{}
}

// GetSoftwareInfo retrieves installed packages and runtimes.
func (p *SoftwareProvider) GetSoftwareInfo(ctx context.Context) (models.SoftwareInfo, error) {
	// TODO: implement pkg software discovery
	return models.SoftwareInfo{
		Packages: []models.SoftwarePackage{},
		Runtimes: []models.RuntimeEnvironment{},
	}, nil
}
