package discovery

import (
	"context"
	"fmt"
	"time"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// Stage implements discovery.Stage for OS discovery.
type OSStage struct{}

func (s *OSStage) Name() string {
	return "os"
}

func (s *OSStage) Version() string {
	return "1.0.0"
}

func (s *OSStage) Description() string {
	return "Discovers immutable operating system facts using the Platform abstraction layer"
}

func (s *OSStage) Priority() int {
	return 20 // Runs after hardware discovery
}

func (s *OSStage) DependsOn() []string {
	// While it relies on platform detector internally, from a pipeline perspective
	// it might just need hardware context in future.
	return []string{"hardware"}
}

func (s *OSStage) Timeout() time.Duration {
	return 15 * time.Second
}

func (s *OSStage) Initialize(dctx Context) error {
	if dctx.Platform() == nil {
		return fmt.Errorf("platform abstraction layer is not initialized in context")
	}
	if dctx.Platform().OS() == nil {
		return fmt.Errorf("platform OS provider is not initialized")
	}
	return nil
}

func (s *OSStage) Run(ctx context.Context, dctx Context) (DiscoveryArtifact, error) {
	// The stage uses the Platform abstraction exclusively.
	// Zero runtime.GOOS checks exist here.
	provider := dctx.Platform().OS()

	info, err := provider.GetOSInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to discover OS info: %w", err)
	}

	return info, nil
}

func (s *OSStage) Validate(artifact DiscoveryArtifact) error {
	info, ok := artifact.(models.OSInfo)
	if !ok {
		return fmt.Errorf("expected models.OSInfo artifact, got %T", artifact)
	}

	if info.OperatingSystem == "" {
		return fmt.Errorf("invalid artifact: missing operating system")
	}

	return nil
}

func (s *OSStage) Cleanup(ctx context.Context) error {
	return nil
}
