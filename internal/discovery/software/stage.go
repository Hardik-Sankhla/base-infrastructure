package software

import (
	"context"
	"fmt"

	"github.com/base-infrastructure/platform/internal/discovery"
)

// Stage discovers installed software and runtimes.
type Stage struct{}

// NewStage creates a new software discovery stage.
func NewStage() *Stage {
	return &Stage{}
}

// Name returns the name of the stage.
func (s *Stage) Name() string {
	return "software"
}

// Dependencies returns the dependencies of the stage.
func (s *Stage) Dependencies() []string {
	return []string{}
}

// Initialize prepares the stage for execution.
func (s *Stage) Initialize(ctx discovery.Context) error {
	if ctx.Platform() == nil {
		return fmt.Errorf("platform not found in context")
	}
	return nil
}

// Run executes the discovery process.
func (s *Stage) Run(ctx context.Context, dctx discovery.Context) (interface{}, error) {
	dctx.Logger().Debug("starting software discovery")

	provider := dctx.Platform().Software()
	info, err := provider.GetSoftwareInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get software info: %w", err)
	}

	dctx.Logger().Debug(
		"completed software discovery",
		"packages_count", len(info.Packages),
		"runtimes_count", len(info.Runtimes),
	)

	return info, nil
}

// Validate ensures the artifact is valid.
func (s *Stage) Validate(artifact interface{}) error {
	if artifact == nil {
		return fmt.Errorf("software artifact is nil")
	}
	return nil
}
