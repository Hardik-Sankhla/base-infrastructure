package discovery

import (
	"context"
	"fmt"
	"time"
)

// Stage implements discovery.Stage for filesystem discovery.
type FilesystemStage struct{}

func (s *FilesystemStage) Name() string {
	return "filesystem"
}

func (s *FilesystemStage) Version() string {
	return "1.0.0"
}

func (s *FilesystemStage) Description() string {
	return "Discovers filesystem mounts, capacities, standard directories, and capabilities"
}

func (s *FilesystemStage) Priority() int {
	return 30 // Runs after OS (20)
}

func (s *FilesystemStage) DependsOn() []string {
	return []string{"os"}
}

func (s *FilesystemStage) Timeout() time.Duration {
	return 10 * time.Second
}

func (s *FilesystemStage) Initialize(dctx Context) error {
	if dctx.Platform() == nil {
		return fmt.Errorf("platform abstraction layer is not initialized in context")
	}
	if dctx.Platform().Filesystem() == nil {
		return fmt.Errorf("filesystem provider is not available for this platform")
	}
	return nil
}

func (s *FilesystemStage) Run(ctx context.Context, dctx Context) (DiscoveryArtifact, error) {
	provider := dctx.Platform().Filesystem()

	info, err := provider.GetFilesystemInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to discover filesystem info: %w", err)
	}

	return info, nil
}

func (s *FilesystemStage) Validate(artifact DiscoveryArtifact) error {
	if artifact == nil {
		return fmt.Errorf("artifact is nil")
	}
	if artifact.ArtifactType() != "Filesystem" {
		return fmt.Errorf("expected Filesystem artifact, got %s", artifact.ArtifactType())
	}
	return nil
}

func (s *FilesystemStage) Cleanup(ctx context.Context) error {
	return nil
}
