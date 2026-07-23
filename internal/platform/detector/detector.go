package detector

import (
	"github.com/base-infrastructure/platform/internal/platform"
)

// DefaultDetector implements platform.Detector using standard Go runtime checks
// and heuristic fallbacks (e.g., to differentiate WSL from generic Linux).
type DefaultDetector struct{}

// NewDetector returns a DefaultDetector.
func NewDetector() *DefaultDetector {
	return &DefaultDetector{}
}

// Detect determines the current platform and returns the appropriate abstraction.
func (d *DefaultDetector) Detect() (platform.Platform, error) {
	return platform.NewPlatform(), nil
}
