package contracts

import (
	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/runtime/context"
)

// DiscoveryResult is the output of the Discovery Engine
type DiscoveryResult struct {
	Hardware     models.Hardware
	OS           models.OSInfo
	Environment  models.Environment
	Capabilities []models.Capability
}

// DiscoveryEngine Contract
type DiscoveryEngine interface {
	Run(ctx *context.PlatformContext) (DiscoveryResult, error)
}

// PlannerEngine Contract
type PlannerEngine interface {
	Plan(ctx *context.PlatformContext, discovery DiscoveryResult, policies models.Policy) (models.ExecutionPlan, error)
}

// ExecutorEngine Contract
type ExecutorEngine interface {
	Execute(ctx *context.PlatformContext, plan models.ExecutionPlan) (models.Result, error)
}

// ValidatorEngine Contract
type ValidatorEngine interface {
	Validate(ctx *context.PlatformContext, execResult models.Result) (models.Result, error)
}
