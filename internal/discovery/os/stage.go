package os

import (
	"context"
	"fmt"
	"time"

	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/domain/models"
)

// Stage implements discovery.Stage for OS discovery.
type Stage struct{}

// NewStage creates a new OS discovery stage.
func NewStage() *Stage {
	return &Stage{}
}

func (s *Stage) Name() string {
	return "os"
}

func (s *Stage) Version() string {
	return "1.0.0"
}

func (s *Stage) Description() string {
	return "Discovers immutable operating system facts using the Platform abstraction layer"
}

func (s *Stage) Priority() int {
	return 20 // Runs after hardware discovery
}

func (s *Stage) DependsOn() []string {
	// While it relies on platform detector internally, from a pipeline perspective
	// it might just need hardware context in future.
	return []string{"hardware"}
}

func (s *Stage) Timeout() time.Duration {
	return 15 * time.Second
}

func (s *Stage) Initialize(dctx discovery.Context) error {
	if dctx.Platform() == nil {
		return fmt.Errorf("platform abstraction layer is not initialized in context")
	}
	if dctx.Platform().OS() == nil {
		return fmt.Errorf("platform OS provider is not initialized")
	}
	return nil
}

func (s *Stage) Run(ctx context.Context, dctx discovery.Context) (discovery.DiscoveryArtifact, error) {
	// The stage uses the Platform abstraction exclusively.
	// Zero runtime.GOOS checks exist here.
	provider := dctx.Platform().OS()

	info, err := provider.GetOSInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to discover OS info: %w", err)
	}

	return info, nil
}

func (s *Stage) Validate(artifact discovery.DiscoveryArtifact) error {
	info, ok := artifact.(models.OSInfo)
	if !ok {
		return fmt.Errorf("expected models.OSInfo artifact, got %T", artifact)
	}

	if info.OperatingSystem == "" {
		return fmt.Errorf("invalid artifact: missing operating system")
	}

	return nil
}

func (s *Stage) Cleanup(ctx context.Context) error {
	return nil
}
