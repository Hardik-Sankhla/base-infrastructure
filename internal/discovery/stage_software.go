package discovery

import (
	"context"
	"fmt"
	"time"
)

// Stage discovers installed software and runtimes.
type SoftwareStage struct{}

// Name returns the name of the stage.
func (s *SoftwareStage) Name() string {
	return "software"
}

// Version returns the version of the stage.
func (s *SoftwareStage) Version() string {
	return "1.0.0"
}

// Description returns a description of what the stage discovers.
func (s *SoftwareStage) Description() string {
	return "Discovers installed software and runtimes"
}

// Priority returns the priority of the stage (lower runs first).
func (s *SoftwareStage) Priority() int {
	return 60
}

// DependsOn returns the dependencies of the stage.
func (s *SoftwareStage) DependsOn() []string {
	return []string{}
}

// Timeout returns the maximum duration allowed for this stage.
func (s *SoftwareStage) Timeout() time.Duration {
	return 30 * time.Second
}

// Initialize prepares the stage for execution.
func (s *SoftwareStage) Initialize(ctx Context) error {
	if ctx.Platform() == nil {
		return fmt.Errorf("platform not found in context")
	}
	return nil
}

// Run executes the discovery process.
func (s *SoftwareStage) Run(ctx context.Context, dctx Context) (DiscoveryArtifact, error) {
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
func (s *SoftwareStage) Validate(artifact DiscoveryArtifact) error {
	if artifact == nil {
		return fmt.Errorf("software artifact is nil")
	}
	return nil
}

// Cleanup performs any necessary resource cleanup after the stage completes.
func (s *SoftwareStage) Cleanup(ctx context.Context) error {
	return nil
}
