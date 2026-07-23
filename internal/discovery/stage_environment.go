package discovery

import (
	"context"
	"fmt"
	"time"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// Stage implements discovery.Stage for environment discovery.
type EnvironmentStage struct{}

func (s *EnvironmentStage) Name() string {
	return "environment"
}

func (s *EnvironmentStage) Version() string {
	return "1.0.0"
}

func (s *EnvironmentStage) Description() string {
	return "Discovers the execution environment context (VM, Container, Cloud, CI/CD, Session)"
}

func (s *EnvironmentStage) Priority() int {
	return 40 // Runs after Network (30)
}

func (s *EnvironmentStage) DependsOn() []string {
	return []string{"os"} // Requires basic OS facts logically
}

func (s *EnvironmentStage) Timeout() time.Duration {
	return 10 * time.Second
}

func (s *EnvironmentStage) Initialize(dctx Context) error {
	if dctx.Platform() == nil {
		return fmt.Errorf("platform abstraction layer is not initialized in context")
	}
	if dctx.Platform().Environment() == nil {
		return fmt.Errorf("environment provider is not available for this platform")
	}
	return nil
}

func (s *EnvironmentStage) Run(ctx context.Context, dctx Context) (DiscoveryArtifact, error) {
	provider := dctx.Platform().Environment()

	envInfo, err := provider.GetEnvironmentInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to discover environment context: %w", err)
	}

	return envInfo, nil
}

func (s *EnvironmentStage) Validate(artifact DiscoveryArtifact) error {
	_, ok := artifact.(models.EnvironmentInfo)
	if !ok {
		return fmt.Errorf("expected models.EnvironmentInfo artifact, got %T", artifact)
	}

	return nil
}

func (s *EnvironmentStage) Cleanup(ctx context.Context) error {
	// Nothing to clean up
	return nil
}
