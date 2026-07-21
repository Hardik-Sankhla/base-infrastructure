package software

import (
	"context"
	"fmt"
	"time"

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

// Version returns the version of the stage.
func (s *Stage) Version() string {
	return "1.0.0"
}

// Description returns a description of what the stage discovers.
func (s *Stage) Description() string {
	return "Discovers installed software and runtimes"
}

// Priority returns the priority of the stage (lower runs first).
func (s *Stage) Priority() int {
	return 60
}

// DependsOn returns the dependencies of the stage.
func (s *Stage) DependsOn() []string {
	return []string{}
}

// Timeout returns the maximum duration allowed for this stage.
func (s *Stage) Timeout() time.Duration {
	return 30 * time.Second
}

// Initialize prepares the stage for execution.
func (s *Stage) Initialize(ctx discovery.Context) error {
	if ctx.Platform() == nil {
		return fmt.Errorf("platform not found in context")
	}
	return nil
}

// Run executes the discovery process.
func (s *Stage) Run(ctx context.Context, dctx discovery.Context) (discovery.DiscoveryArtifact, error) {
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
func (s *Stage) Validate(artifact discovery.DiscoveryArtifact) error {
	if artifact == nil {
		return fmt.Errorf("software artifact is nil")
	}
	return nil
}

// Cleanup performs any necessary resource cleanup after the stage completes.
func (s *Stage) Cleanup(ctx context.Context) error {
	return nil
}
