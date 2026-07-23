package contracts

import (
	"github.com/base-infrastructure/platform/internal/runtime"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// DiscoveryEngine orchestrates the discovery phase.
type DiscoveryEngine interface {
	// Run executes the discovery pipeline and returns the discovery manifest.
	Run(ctx *runtime.PlatformContext) (*models.DiscoveryManifest, error)
}

// PlannerEngine Contract
type PlannerEngine interface {
	Plan(ctx *runtime.PlatformContext, dr *models.DiscoveryManifest, policies models.Policy) (models.ExecutionPlan, error)
}

// ExecutorEngine Contract
type ExecutorEngine interface {
	Execute(ctx *runtime.PlatformContext, plan models.ExecutionPlan) (models.Result, error)
}

// ValidatorEngine Contract
type ValidatorEngine interface {
	Validate(ctx *runtime.PlatformContext, execResult models.Result) (models.Result, error)
}
