package platform

import (
	"context"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// OSProvider abstracts the retrieval of operating system information.
type OSProvider interface {
	GetOSInfo(ctx context.Context) (models.OSInfo, error)
}

// Platform provides a cross-platform abstraction for discovery and execution.
type Platform interface {
	// ID returns a unique identifier for the platform (e.g. "linux", "windows", "wsl")
	ID() string
	// Name returns a human-readable name of the platform
	Name() string

	// OS returns the OS discovery provider
	OS() OSProvider

	// (Future) Filesystem() FilesystemProvider
	// (Future) Network() NetworkProvider
}

// Detector is responsible for identifying the current runtime environment
// and returning the appropriate Platform implementation.
type Detector interface {
	Detect() (Platform, error)
}
