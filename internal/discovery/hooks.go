package discovery

import (
	"github.com/base-infrastructure/platform/internal/runtime"
)

// Hook represents a lifecycle hook that can be attached to the discovery pipeline.
// Not all methods need to be implemented; implementations can just provide what they need.
type Hook interface {
	// BeforePipeline is called before the pipeline starts executing stages.
	BeforePipeline(ctx runtime.Context, dctx Context, stages []Stage) error

	// AfterPipeline is called after the pipeline has finished (success or failure).
	AfterPipeline(ctx runtime.Context, dctx Context, results map[string]StageResult) error

	// BeforeStage is called right before a specific stage executes.
	BeforeStage(ctx runtime.Context, dctx Context, stage Stage) error

	// AfterStage is called right after a specific stage executes successfully.
	AfterStage(ctx runtime.Context, dctx Context, stage Stage, artifact DiscoveryArtifact) error

	// OnStageError is called when a specific stage fails.
	OnStageError(ctx runtime.Context, dctx Context, stage Stage, err error)
}

// BaseHook provides empty implementations for all Hook methods so implementers
// can embed it and only override what they care about.
type BaseHook struct{}

func (b *BaseHook) BeforePipeline(ctx runtime.Context, dctx Context, stages []Stage) error {
	return nil
}

func (b *BaseHook) AfterPipeline(ctx runtime.Context, dctx Context, results map[string]StageResult) error {
	return nil
}

func (b *BaseHook) BeforeStage(ctx runtime.Context, dctx Context, stage Stage) error {
	return nil
}

func (b *BaseHook) AfterStage(ctx runtime.Context, dctx Context, stage Stage, artifact DiscoveryArtifact) error {
	return nil
}

func (b *BaseHook) OnStageError(ctx runtime.Context, dctx Context, stage Stage, err error) {}
