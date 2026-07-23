package builtin

import (
	"github.com/base-infrastructure/platform/internal/core"
	"github.com/base-infrastructure/platform/internal/discovery"
)

// DefaultStages returns the list of built-in discovery stages in recommended order.
func DefaultStages() []core.Stage {
	return []core.Stage{
		&discovery.HardwareStage{},
		&discovery.OSStage{},
		&discovery.NetworkStage{},
		&discovery.EnvironmentStage{},
		&discovery.FilesystemStage{},
		&discovery.SoftwareStage{},
	}
}

// RegisterCoreStages registers all built-in discovery stages into the provided registry.
func RegisterCoreStages(reg *core.Registry) error {
	// Register Hardware Stage (PR #2)
	if err := reg.Register(&discovery.HardwareStage{}); err != nil {
		return err
	}

	// Register OS Stage (PR #3)
	if err := reg.Register(&discovery.OSStage{}); err != nil {
		return err
	}

	// Register Network Stage (PR #5)
	if err := reg.Register(&discovery.NetworkStage{}); err != nil {
		return err
	}

	// Register Environment Stage (PR #6)
	if err := reg.Register(&discovery.EnvironmentStage{}); err != nil {
		return err
	}

	// Register Filesystem Stage (PR #4)
	if err := reg.Register(&discovery.FilesystemStage{}); err != nil {
		return err
	}

	// Register Software Stage (PR #7)
	if err := reg.Register(&discovery.SoftwareStage{}); err != nil {
		return err
	}

	return nil
}
