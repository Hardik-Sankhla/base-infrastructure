package hardware

import (
	"context"
	"fmt"
	"time"

	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/domain/models"
)

// Stage implements discovery.Stage for hardware discovery.
type Stage struct {
	provider Provider
}

// NewStage creates a new Hardware discovery stage using the given provider.
func NewStage(provider Provider) *Stage {
	return &Stage{
		provider: provider,
	}
}

func (s *Stage) Name() string {
	return "hardware"
}

func (s *Stage) Version() string {
	return "1.0.0"
}

func (s *Stage) Description() string {
	return "Discovers physical hardware resources (CPU, RAM, Storage, GPU, Battery, Thermal)"
}

func (s *Stage) Priority() int {
	return 10 // Runs early in the pipeline
}

func (s *Stage) DependsOn() []string {
	return nil // Base stage, no dependencies
}

func (s *Stage) Timeout() time.Duration {
	return 30 * time.Second
}

func (s *Stage) Initialize(dctx discovery.Context) error {
	if s.provider == nil {
		return fmt.Errorf("hardware provider is not initialized")
	}
	return nil
}

func (s *Stage) Run(ctx context.Context, dctx discovery.Context) (discovery.DiscoveryArtifact, error) {
	var hw models.Hardware
	var err error

	// Critical components (must succeed)
	if hw.CPU, err = s.provider.GetCPU(ctx); err != nil {
		return nil, fmt.Errorf("failed to discover CPU: %w", err)
	}

	if hw.RAM, err = s.provider.GetRAM(ctx); err != nil {
		return nil, fmt.Errorf("failed to discover RAM: %w", err)
	}

	if hw.Storage, err = s.provider.GetStorage(ctx); err != nil {
		return nil, fmt.Errorf("failed to discover Storage: %w", err)
	}

	// Non-critical components (graceful fallback)
	if gpus, gErr := s.provider.GetGPUs(ctx); gErr == nil {
		hw.GPUs = gpus
	} else {
		dctx.Logger().Debug("Failed to discover GPUs or none present", "error", gErr)
	}

	if battery, bErr := s.provider.GetBattery(ctx); bErr == nil {
		hw.Battery = battery
	} else {
		dctx.Logger().Debug("Failed to discover Battery or none present", "error", bErr)
	}

	if thermals, tErr := s.provider.GetThermal(ctx); tErr == nil {
		hw.Thermals = thermals
	} else {
		dctx.Logger().Debug("Failed to discover Thermal sensors or none present", "error", tErr)
	}

	return hw, nil
}

func (s *Stage) Validate(artifact discovery.DiscoveryArtifact) error {
	hw, ok := artifact.(models.Hardware)
	if !ok {
		return fmt.Errorf("expected models.Hardware artifact, got %T", artifact)
	}

	if hw.CPU.Architecture == "" {
		return fmt.Errorf("invalid artifact: missing CPU architecture")
	}
	if hw.RAM.TotalBytes == 0 {
		return fmt.Errorf("invalid artifact: missing RAM total bytes")
	}
	if len(hw.Storage) == 0 {
		return fmt.Errorf("invalid artifact: missing storage devices")
	}

	return nil
}

func (s *Stage) Cleanup(ctx context.Context) error {
	// Nothing to clean up for hardware discovery
	return nil
}
