package discovery

import (
	"context"
)

// Hook represents a lifecycle hook that can be attached to the discovery pipeline.
// Not all methods need to be implemented; implementations can just provide what they need.
type Hook interface {
	// BeforePipeline is called before the pipeline starts executing stages.
	BeforePipeline(ctx context.Context, dctx Context, stages []Stage) error

	// AfterPipeline is called after the pipeline has finished (success or failure).
	AfterPipeline(ctx context.Context, dctx Context, results map[string]StageResult) error

	// BeforeStage is called right before a specific stage executes.
	BeforeStage(ctx context.Context, dctx Context, stage Stage) error

	// AfterStage is called right after a specific stage executes successfully.
	AfterStage(ctx context.Context, dctx Context, stage Stage, artifact DiscoveryArtifact) error

	// OnStageError is called when a specific stage fails.
	OnStageError(ctx context.Context, dctx Context, stage Stage, err error)
}

// BaseHook provides empty implementations for all Hook methods so implementers
// can embed it and only override what they care about.
type BaseHook struct{}

func (b *BaseHook) BeforePipeline(ctx context.Context, dctx Context, stages []Stage) error {
	return nil
}

func (b *BaseHook) AfterPipeline(ctx context.Context, dctx Context, results map[string]StageResult) error {
	return nil
}

func (b *BaseHook) BeforeStage(ctx context.Context, dctx Context, stage Stage) error {
	return nil
}

func (b *BaseHook) AfterStage(ctx context.Context, dctx Context, stage Stage, artifact DiscoveryArtifact) error {
	return nil
}

func (b *BaseHook) OnStageError(ctx context.Context, dctx Context, stage Stage, err error) {}
