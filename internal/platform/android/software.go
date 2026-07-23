package android

import (
	"context"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// SoftwareProvider implements platform.SoftwareProvider for Android.
type SoftwareProvider struct{}

// NewSoftwareProvider creates a new Android software provider.
func NewSoftwareProvider() *SoftwareProvider {
	return &SoftwareProvider{}
}

// GetSoftwareInfo retrieves installed packages and runtimes.
func (p *SoftwareProvider) GetSoftwareInfo(ctx context.Context) (models.SoftwareInfo, error) {
	// TODO: implement termux/android software discovery
	return models.SoftwareInfo{
		Packages: []models.SoftwarePackage{},
		Runtimes: []models.RuntimeEnvironment{},
	}, nil
}
