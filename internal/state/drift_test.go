package state

import (
	"testing"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

func TestDriftDetector_Detect(t *testing.T) {
	detector := NewDriftDetector()

	tests := []struct {
		name     string
		desired  models.StateManifest
		current  []models.Capability
		wantSize int
		wantType models.DriftType
	}{
		{
			name: "No drift when exact match",
			desired: models.StateManifest{
				Capabilities: []models.DesiredCapability{
					{ID: "docker", State: models.StateAvailable, Version: "20.10.0"},
				},
			},
			current: []models.Capability{
				{ID: "docker", State: models.StateAvailable, Version: "20.10.0"},
			},
			wantSize: 0,
		},
		{
			name: "Missing capability",
			desired: models.StateManifest{
				Capabilities: []models.DesiredCapability{
					{ID: "docker", State: models.StateAvailable},
				},
			},
			current:  []models.Capability{},
			wantSize: 1,
			wantType: models.DriftMissing,
		},
		{
			name: "State mismatch",
			desired: models.StateManifest{
				Capabilities: []models.DesiredCapability{
					{ID: "docker", State: models.StateAvailable},
				},
			},
			current: []models.Capability{
				{ID: "docker", State: models.StateMissing},
			},
			wantSize: 1,
			wantType: models.DriftMissing,
		},
		{
			name: "Strict version mismatch",
			desired: models.StateManifest{
				Settings: models.StateSettings{
					StrictVersionMatch: true,
				},
				Capabilities: []models.DesiredCapability{
					{ID: "docker", State: models.StateAvailable, Version: "24.0.0"},
				},
			},
			current: []models.Capability{
				{ID: "docker", State: models.StateAvailable, Version: "20.10.0"},
			},
			wantSize: 1,
			wantType: models.DriftVersion,
		},
		{
			name: "Lenient version match skips strict check",
			desired: models.StateManifest{
				Settings: models.StateSettings{
					StrictVersionMatch: false,
				},
				Capabilities: []models.DesiredCapability{
					{ID: "docker", State: models.StateAvailable, Version: "24.0.0"},
				},
			},
			current: []models.Capability{
				{ID: "docker", State: models.StateAvailable, Version: "20.10.0"},
			},
			wantSize: 0, // Not implemented fully yet, so it skips
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			drifts, err := detector.Detect(tt.desired, tt.current)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(drifts) != tt.wantSize {
				t.Errorf("got %d drifts, want %d", len(drifts), tt.wantSize)
			}
			if tt.wantSize > 0 && drifts[0].Type != tt.wantType {
				t.Errorf("got drift type %s, want %s", drifts[0].Type, tt.wantType)
			}
		})
	}
}
