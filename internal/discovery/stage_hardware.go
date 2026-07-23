package discovery

import (
	"context"
	"fmt"
	"time"

	"github.com/base-infrastructure/platform/internal/core"
	"github.com/base-infrastructure/platform/internal/domain/models"
)

// core.Stage implements discovery.core.Stage for hardware discovery.
type HardwareStage struct{}

func (s *HardwareStage) Name() string {
	return "hardware"
}

func (s *HardwareStage) Version() string {
	return "1.0.0"
}

func (s *HardwareStage) Description() string {
	return "Discovers physical hardware resources (CPU, RAM, Storage, GPU, Battery, Thermal)"
}

func (s *HardwareStage) Priority() int {
	return 10 // Runs early in the pipeline
}

func (s *HardwareStage) DependsOn() []string {
	return nil // Base stage, no dependencies
}

func (s *HardwareStage) Timeout() time.Duration {
	return 30 * time.Second
}

func (s *HardwareStage) Initialize(dctx core.Context) error {
	if dctx.Platform() == nil {
		return fmt.Errorf("platform abstraction layer is not initialized in context")
	}
	if dctx.Platform().Hardware() == nil {
		return fmt.Errorf("hardware provider is not available for this platform")
	}
	return nil
}

func (s *HardwareStage) Run(ctx context.Context, dctx core.Context) (core.DiscoveryArtifact, error) {
	var hw models.Hardware
	var err error

	provider := dctx.Platform().Hardware()

	// Critical components (must succeed)
	if hw.CPU, err = provider.GetCPU(ctx); err != nil {
		return nil, fmt.Errorf("failed to discover CPU: %w", err)
	}

	if hw.RAM, err = provider.GetRAM(ctx); err != nil {
		return nil, fmt.Errorf("failed to discover RAM: %w", err)
	}

	if hw.Storage, err = provider.GetStorage(ctx); err != nil {
		return nil, fmt.Errorf("failed to discover Storage: %w", err)
	}

	// Non-critical components (graceful fallback)
	if gpus, gErr := provider.GetGPUs(ctx); gErr == nil {
		hw.GPUs = gpus
	} else {
		dctx.Logger().Debug("Failed to discover GPUs or none present", "error", gErr)
	}

	if battery, bErr := provider.GetBattery(ctx); bErr == nil {
		hw.Battery = battery
	} else {
		dctx.Logger().Debug("Failed to discover Battery or none present", "error", bErr)
	}

	if thermals, tErr := provider.GetThermal(ctx); tErr == nil {
		hw.Thermals = thermals
	} else {
		dctx.Logger().Debug("Failed to discover Thermal sensors or none present", "error", tErr)
	}

	return hw, nil
}

func (s *HardwareStage) Validate(artifact core.DiscoveryArtifact) error {
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

func (s *HardwareStage) Cleanup(ctx context.Context) error {
	// Nothing to clean up for hardware discovery
	return nil
}
