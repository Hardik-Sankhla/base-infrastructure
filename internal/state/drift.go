package state

import (
	"github.com/base-infrastructure/platform/internal/domain/models"
)

// DriftDetector compares desired state against current capabilities
type DriftDetector interface {
	Detect(desired models.StateManifest, current []models.Capability) ([]models.Drift, error)
}

type driftDetector struct{}

// NewDriftDetector creates a new drift detector
func NewDriftDetector() DriftDetector {
	return &driftDetector{}
}

// Detect computes the drift between desired and current state
func (d *driftDetector) Detect(desired models.StateManifest, current []models.Capability) ([]models.Drift, error) {
	var drifts []models.Drift

	// Map current capabilities for O(1) lookup
	currentMap := make(map[string]models.Capability)
	for _, cap := range current {
		currentMap[cap.ID] = cap
	}

	// Detect missing or mismatched desired capabilities
	for _, req := range desired.Capabilities {
		curr, exists := currentMap[req.ID]

		if !exists {
			// Capability is completely missing
			drifts = append(drifts, models.Drift{
				CapabilityID: req.ID,
				Type:         models.DriftMissing,
				Desired:      req,
				Current:      nil,
			})
			continue
		}

		// Check state
		if req.State != "" && req.State != curr.State {
			drifts = append(drifts, models.Drift{
				CapabilityID: req.ID,
				Type:         models.DriftMissing, // Conceptually, the desired state is missing
				Desired:      req,
				Current:      &curr,
			})
			continue
		}

		// Check version (naive exact match for now)
		if req.Version != "" && curr.Version != "" && req.Version != curr.Version {
			if desired.Settings.StrictVersionMatch {
				drifts = append(drifts, models.Drift{
					CapabilityID: req.ID,
					Type:         models.DriftVersion,
					Desired:      req,
					Current:      &curr,
				})
				continue
			}
			// TODO: Implement semver parsing if strict match is false
		}
	}

	return drifts, nil
}
