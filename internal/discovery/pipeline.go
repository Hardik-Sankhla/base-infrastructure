package discovery

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

// PipelineConfig controls pipeline execution behaviour.
type PipelineConfig struct {
	FailFast bool
}

// DefaultPipelineConfig returns a PipelineConfig with sensible defaults.
func DefaultPipelineConfig() PipelineConfig {
	return PipelineConfig{
		FailFast: true,
	}
}

// Pipeline executes an ordered sequence of discovery stages.
type Pipeline struct {
	stages []Stage
	hooks  []Hook
	config PipelineConfig
	logger *slog.Logger
	bus    events.EventBus
}

// NewPipeline creates a pipeline with the given configuration.
func NewPipeline(cfg PipelineConfig, logger *slog.Logger, bus events.EventBus) *Pipeline {
	return &Pipeline{
		stages: make([]Stage, 0),
		hooks:  make([]Hook, 0),
		config: cfg,
		logger: logger,
		bus:    bus,
	}
}

// AddHook registers a pipeline hook.
func (p *Pipeline) AddHook(h Hook) {
	if h != nil {
		p.hooks = append(p.hooks, h)
	}
}

// AddStage appends a stage to the pipeline. Stages are re-sorted by
// priority before each Run.
func (p *Pipeline) AddStage(stage Stage) error {
	if stage == nil {
		return fmt.Errorf("cannot add nil stage to pipeline")
	}
	for _, s := range p.stages {
		if s.Name() == stage.Name() {
			return fmt.Errorf("stage %q already exists in pipeline", stage.Name())
		}
	}
	p.stages = append(p.stages, stage)
	return nil
}

// AddStages is a convenience method that adds multiple stages at once.
func (p *Pipeline) AddStages(stages []Stage) error {
	for _, s := range stages {
		if err := p.AddStage(s); err != nil {
			return err
		}
	}
	return nil
}

// Run executes all stages sequentially in priority order.
func (p *Pipeline) Run(ctx context.Context, dctx Context) (*Result, error) {
	// Validate the dependency graph before execution
	validator := NewValidator()
	if err := validator.Validate(p.stages); err != nil {
		return nil, fmt.Errorf("pipeline validation failed: %w", err)
	}

	builder := NewResultBuilder()
	sorted := p.sortedStages()
	builder.SetTotalStages(len(sorted))

	p.logger.Info("Discovery pipeline starting", "stages", len(sorted))

	// Hook: BeforePipeline
	for _, h := range p.hooks {
		if err := h.BeforePipeline(ctx, dctx, sorted); err != nil {
			return nil, fmt.Errorf("BeforePipeline hook failed: %w", err)
		}
	}

	for _, stage := range sorted {
		if err := ctx.Err(); err != nil {
			p.logger.Warn("Pipeline cancelled", "reason", err)
			builder.AddStageError("_pipeline", fmt.Errorf("pipeline cancelled: %w", err))
			break
		}

		result, err := p.runStage(ctx, dctx, stage)
		builder.AddStageResult(result)

		if err != nil {
			if p.config.FailFast {
				p.logger.Error("Pipeline aborting (fail-fast)", "stage", stage.Name(), "error", err)
				break
			}
			p.logger.Warn("Stage failed, continuing", "stage", stage.Name(), "error", err)
			continue
		}
	}

	finalResult := builder.Build()
	p.logger.Info(
		"Discovery pipeline finished",
		"success", finalResult.Success,
		"duration", finalResult.Duration,
		"total", finalResult.TotalStages,
		"successful", finalResult.SuccessfulStages,
		"failed", finalResult.FailedStages,
		"skipped", finalResult.SkippedStages,
	)

	// Hook: AfterPipeline
	stageResults := make(map[string]StageResult)
	for name, res := range finalResult.Stages {
		stageResults[name] = *res
	}
	for _, h := range p.hooks {
		if err := h.AfterPipeline(ctx, dctx, stageResults); err != nil {
			p.logger.Warn("AfterPipeline hook failed", "error", err)
		}
	}

	return finalResult, nil
}

// runStage executes a single stage with timing, logging, event publishing,
// and the full stage lifecycle (Init, Run, Validate, Cleanup).
func (p *Pipeline) runStage(globalCtx context.Context, dctx Context, stage Stage) (*StageResult, error) {
	name := stage.Name()
	p.logger.Info("Stage starting", "stage", name)

	// Combine global context with stage timeout.
	var ctx context.Context
	var cancel context.CancelFunc
	if timeout := stage.Timeout(); timeout > 0 {
		ctx, cancel = context.WithTimeout(globalCtx, timeout)
	} else {
		// Use a default reasonable timeout to prevent hanging stages
		ctx, cancel = context.WithTimeout(globalCtx, 5*time.Minute)
	}
	defer cancel()

	p.publishEvent(events.DiscoveryStageStarted, events.StageEventPayload{
		StageName: name,
		Status:    string(StatusRunning),
		Timestamp: time.Now(),
	})

	start := time.Now()
	var artifact DiscoveryArtifact
	var err error

	// Defer cleanup to ensure it runs even on panic or error
	defer func() {
		if cleanupErr := stage.Cleanup(context.Background()); cleanupErr != nil {
			p.logger.Warn("Stage cleanup failed", "stage", name, "error", cleanupErr)
		}
	}()

	// Hook: BeforeStage
	for _, h := range p.hooks {
		if err = h.BeforeStage(ctx, dctx, stage); err != nil {
			err = fmt.Errorf("BeforeStage hook failed: %w", err)
			goto HandleError
		}
	}

	// Lifecycle execution
	if err = stage.Initialize(dctx); err != nil {
		err = fmt.Errorf("initialize failed: %w", err)
		goto HandleError
	}

	if artifact, err = stage.Run(ctx, dctx); err != nil {
		err = fmt.Errorf("run failed: %w", err)
		goto HandleError
	}

	if err = stage.Validate(artifact); err != nil {
		err = fmt.Errorf("validation failed: %w", err)
		goto HandleError
	}

	{
		// Success Path
		elapsed := time.Since(start)
		p.logger.Info("Stage completed", "stage", name, "duration", elapsed)
		p.publishEvent(events.DiscoveryStageCompleted, events.StageEventPayload{
			StageName: name,
			Status:    string(StatusSuccess),
			Timestamp: time.Now(),
			Duration:  elapsed,
		})

		// Hook: AfterStage
		for _, h := range p.hooks {
			if hErr := h.AfterStage(ctx, dctx, stage, artifact); hErr != nil {
				p.logger.Warn("AfterStage hook failed", "stage", name, "error", hErr)
			}
		}

		return &StageResult{
			StageName: name,
			Status:    StatusSuccess,
			Artifact:  artifact,
			Duration:  elapsed,
			Timestamp: time.Now().UTC(),
		}, nil
	}

HandleError:
	elapsed := time.Since(start)
	p.logger.Error("Stage failed", "stage", name, "duration", elapsed, "error", err)
	p.publishEvent(events.DiscoveryStageFailed, events.StageEventPayload{
		StageName: name,
		Status:    string(StatusFailed),
		Timestamp: time.Now(),
		Duration:  elapsed,
		Error:     err.Error(),
	})

	// Hook: OnStageError
	for _, h := range p.hooks {
		h.OnStageError(ctx, dctx, stage, err)
	}

	return &StageResult{
		StageName: name,
		Status:    StatusFailed,
		Error:     err.Error(),
		Duration:  elapsed,
		Timestamp: time.Now().UTC(),
	}, err
}

func (p *Pipeline) sortedStages() []Stage {
	sorted := make([]Stage, len(p.stages))
	copy(sorted, p.stages)

	for i := 1; i < len(sorted); i++ {
		for j := i; j > 0; j-- {
			if sorted[j].Priority() < sorted[j-1].Priority() {
				sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			} else if sorted[j].Priority() == sorted[j-1].Priority() && sorted[j].Name() < sorted[j-1].Name() {
				sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			} else {
				break
			}
		}
	}
	return sorted
}

func (p *Pipeline) publishEvent(eventType events.EventType, payload any) {
	if p.bus != nil {
		p.bus.Publish(eventType, payload)
	}
}
