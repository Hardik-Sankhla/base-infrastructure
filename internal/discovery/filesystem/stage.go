package filesystem

import (
	"fmt"
	"time"

	"github.com/base-infrastructure/platform/internal/runtime"

	"github.com/base-infrastructure/platform/internal/discovery"
)

// Stage implements discovery.Stage for filesystem discovery.
type Stage struct{}

// NewStage creates a new Filesystem discovery stage.
func NewStage() *Stage {
	return &Stage{}
}

func (s *Stage) Name() string {
	return "filesystem"
}

func (s *Stage) Version() string {
	return "1.0.0"
}

func (s *Stage) Description() string {
	return "Discovers filesystem mounts, capacities, standard directories, and capabilities"
}

func (s *Stage) Priority() int {
	return 30 // Runs after OS (20)
}

func (s *Stage) DependsOn() []string {
	return []string{"os"}
}

func (s *Stage) Timeout() time.Duration {
	return 10 * time.Second
}

func (s *Stage) Initialize(dctx discovery.Context) error {
	if dctx.Platform() == nil {
		return fmt.Errorf("platform abstraction layer is not initialized in context")
	}
	if dctx.Platform().Filesystem() == nil {
		return fmt.Errorf("filesystem provider is not available for this platform")
	}
	return nil
}

func (s *Stage) Run(ctx runtime.Context, dctx discovery.Context) (discovery.DiscoveryArtifact, error) {
	provider := dctx.Platform().Filesystem()

	info, err := provider.GetFilesystemInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to discover filesystem info: %w", err)
	}

	return info, nil
}

func (s *Stage) Validate(artifact discovery.DiscoveryArtifact) error {
	if artifact == nil {
		return fmt.Errorf("artifact is nil")
	}
	if artifact.ArtifactType() != "Filesystem" {
		return fmt.Errorf("expected Filesystem artifact, got %s", artifact.ArtifactType())
	}
	return nil
}

func (s *Stage) Cleanup(ctx runtime.Context) error {
	return nil
}
