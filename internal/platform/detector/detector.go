package detector

import (
	"fmt"
	"runtime"

	"github.com/base-infrastructure/platform/internal/platform"
	"github.com/base-infrastructure/platform/internal/platform/android"
	"github.com/base-infrastructure/platform/internal/platform/bsd"
	"github.com/base-infrastructure/platform/internal/platform/darwin"
	"github.com/base-infrastructure/platform/internal/platform/linux"
	"github.com/base-infrastructure/platform/internal/platform/windows"
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
	// The detector is the ONLY place where runtime.GOOS is allowed.
	// All discovery stages must consume the Platform interface.

	switch runtime.GOOS {
	case "windows":
		return windows.NewPlatform(), nil
	case "linux":
		// Android runtime check usually also reports as linux.
		// We could use heuristics to return android/wsl here instead.
		// For now, we return linux.
		return linux.NewPlatform(), nil
	case "darwin":
		return darwin.NewPlatform(), nil
	case "freebsd", "openbsd", "netbsd":
		return bsd.NewPlatform(), nil
	case "android":
		return android.NewPlatform(), nil
	default:
		return nil, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}
