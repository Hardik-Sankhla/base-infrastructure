package contracts

import (
	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/runtime/context"
)

// DiscoveryEngine orchestrates the discovery phase.
type DiscoveryEngine interface {
	// Run executes the discovery pipeline and returns the discovery manifest.
	Run(ctx *context.PlatformContext) (*models.DiscoveryManifest, error)
}

// PlannerEngine Contract
type PlannerEngine interface {
	Plan(ctx *context.PlatformContext, dr *models.DiscoveryManifest, policies models.Policy) (models.ExecutionPlan, error)
}

// ExecutorEngine Contract
type ExecutorEngine interface {
	Execute(ctx *context.PlatformContext, plan models.ExecutionPlan) (models.Result, error)
}

// ValidatorEngine Contract
type ValidatorEngine interface {
	Validate(ctx *context.PlatformContext, execResult models.Result) (models.Result, error)
}
