package environment

import (
	"github.com/base-infrastructure/platform/internal/runtime"

	"fmt"
	"time"

	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/domain/models"
)

// Stage implements discovery.Stage for environment discovery.
type Stage struct{}

// NewStage creates a new Environment discovery stage.
func NewStage() *Stage {
	return &Stage{}
}

func (s *Stage) Name() string {
	return "environment"
}

func (s *Stage) Version() string {
	return "1.0.0"
}

func (s *Stage) Description() string {
	return "Discovers the execution environment context (VM, Container, Cloud, CI/CD, Session)"
}

func (s *Stage) Priority() int {
	return 40 // Runs after Network (30)
}

func (s *Stage) DependsOn() []string {
	return []string{"os"} // Requires basic OS facts logically
}

func (s *Stage) Timeout() time.Duration {
	return 10 * time.Second
}

func (s *Stage) Initialize(dctx discovery.Context) error {
	if dctx.Platform() == nil {
		return fmt.Errorf("platform abstraction layer is not initialized in context")
	}
	if dctx.Platform().Environment() == nil {
		return fmt.Errorf("environment provider is not available for this platform")
	}
	return nil
}

func (s *Stage) Run(ctx runtime.Context, dctx discovery.Context) (discovery.DiscoveryArtifact, error) {
	provider := dctx.Platform().Environment()

	envInfo, err := provider.GetEnvironmentInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to discover environment context: %w", err)
	}

	return envInfo, nil
}

func (s *Stage) Validate(artifact discovery.DiscoveryArtifact) error {
	_, ok := artifact.(models.EnvironmentInfo)
	if !ok {
		return fmt.Errorf("expected models.EnvironmentInfo artifact, got %T", artifact)
	}

	return nil
}

func (s *Stage) Cleanup(ctx runtime.Context) error {
	// Nothing to clean up
	return nil
}
