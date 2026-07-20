package contracts

import (
	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/runtime/context"
)

// DiscoveryEngine Contract
type DiscoveryEngine interface {
	Run(ctx *context.PlatformContext) (*discovery.Result, error)
}

// PlannerEngine Contract
type PlannerEngine interface {
	Plan(ctx *context.PlatformContext, dr *discovery.Result, policies models.Policy) (models.ExecutionPlan, error)
}

// ExecutorEngine Contract
type ExecutorEngine interface {
	Execute(ctx *context.PlatformContext, plan models.ExecutionPlan) (models.Result, error)
}

// ValidatorEngine Contract
type ValidatorEngine interface {
	Validate(ctx *context.PlatformContext, execResult models.Result) (models.Result, error)
}
