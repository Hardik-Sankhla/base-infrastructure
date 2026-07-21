package builtin

import (
	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/discovery/environment"
	"github.com/base-infrastructure/platform/internal/discovery/filesystem"
	"github.com/base-infrastructure/platform/internal/discovery/hardware"
	"github.com/base-infrastructure/platform/internal/discovery/network"
	"github.com/base-infrastructure/platform/internal/discovery/os"
)

// DefaultStages returns the list of built-in discovery stages in recommended order.
func DefaultStages() []discovery.Stage {
	return []discovery.Stage{
		hardware.NewStage(),
		os.NewStage(),
		network.NewStage(),
		environment.NewStage(),
		filesystem.NewStage(),
	}
}

// RegisterCoreStages registers all built-in discovery stages into the provided registry.
func RegisterCoreStages(reg *discovery.Registry) error {
	// Register Hardware Stage (PR #2)
	if err := reg.Register(hardware.NewStage()); err != nil {
		return err
	}

	// Register OS Stage (PR #3)
	if err := reg.Register(os.NewStage()); err != nil {
		return err
	}

	// Register Network Stage (PR #5)
	if err := reg.Register(network.NewStage()); err != nil {
		return err
	}

	// Register Environment Stage (PR #6)
	if err := reg.Register(environment.NewStage()); err != nil {
		return err
	}

	// Register Filesystem Stage (PR #4)
	if err := reg.Register(filesystem.NewStage()); err != nil {
		return err
	}

	return nil
}
