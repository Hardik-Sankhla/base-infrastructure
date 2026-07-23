package state

import (
	"testing"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

func TestValidator_Validate(t *testing.T) {
	tests := []struct {
		name     string
		manifest *models.StateManifest
		wantErr  bool
	}{
		{
			name: "Valid Manifest",
			manifest: &models.StateManifest{
				Version: "1.0",
				Capabilities: []models.DesiredCapability{
					{
						ID:    "docker",
						State: models.StateAvailable,
					},
					{
						ID:    "legacy-tool",
						State: models.StateMissing,
					},
				},
			},
			wantErr: false,
		},
		{
			name:     "Nil Manifest",
			manifest: nil,
			wantErr:  true,
		},
		{
			name: "Missing Version",
			manifest: &models.StateManifest{
				Version: "",
				Capabilities: []models.DesiredCapability{
					{
						ID:    "docker",
						State: models.StateAvailable,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "No Capabilities",
			manifest: &models.StateManifest{
				Version:      "1.0",
				Capabilities: []models.DesiredCapability{},
			},
			wantErr: true,
		},
		{
			name: "Missing Capability ID",
			manifest: &models.StateManifest{
				Version: "1.0",
				Capabilities: []models.DesiredCapability{
					{
						State: models.StateAvailable,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid Capability State",
			manifest: &models.StateManifest{
				Version: "1.0",
				Capabilities: []models.DesiredCapability{
					{
						ID:    "docker",
						State: "broken", // invalid desired state
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Duplicate Capability ID",
			manifest: &models.StateManifest{
				Version: "1.0",
				Capabilities: []models.DesiredCapability{
					{
						ID:    "docker",
						State: models.StateAvailable,
					},
					{
						ID:    "docker",
						State: models.StateMissing,
					},
				},
			},
			wantErr: true,
		},
	}

	validator := NewValidator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(tt.manifest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
