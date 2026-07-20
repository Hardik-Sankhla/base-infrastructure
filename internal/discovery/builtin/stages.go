package builtin

import (
	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/discovery/filesystem"
	"github.com/base-infrastructure/platform/internal/discovery/hardware"
	"github.com/base-infrastructure/platform/internal/discovery/os"
)

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

	// Register Filesystem Stage (PR #4)
	if err := reg.Register(filesystem.NewStage()); err != nil {
		return err
	}

	return nil
}
