package sdk

import (
	"context"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// Plugin defines the public interface that all 3rd-party plugins must implement.
type Plugin interface {
	// Metadata returns information about this plugin
	Metadata() Metadata

	// Discover probes the local system to see if the target software exists
	Discover(ctx context.Context) (models.Result, error)

	// Install performs the installation payload
	Install(ctx context.Context) (models.Result, error)

	// Verify confirms the installation was successful
	Verify(ctx context.Context) (models.Result, error)

	// Rollback undoes the installation if something failed
	Rollback(ctx context.Context) (models.Result, error)
	
	// Uninstall removes the target completely
	Uninstall(ctx context.Context) (models.Result, error)
}

type Metadata struct {
	Name        string
	Version     string
	Description string
	Provider    string
}
