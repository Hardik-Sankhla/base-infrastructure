package discovery

import (
	"log/slog"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform/detector"
	"github.com/base-infrastructure/platform/internal/runtime/context"
	"github.com/base-infrastructure/platform/internal/runtime/events"
)

// DefaultDiscoveryEngine implements contracts.DiscoveryEngine.
type DefaultDiscoveryEngine struct {
	registry *Registry
	config   PipelineConfig
	logger   *slog.Logger
}

// NewDiscoveryEngine creates a DefaultDiscoveryEngine.
func NewDiscoveryEngine(registry *Registry, cfg PipelineConfig) *DefaultDiscoveryEngine {
	return &DefaultDiscoveryEngine{
		registry: registry,
		config:   cfg,
		logger:   slog.Default(),
	}
}

// Run implements contracts.DiscoveryEngine.
func (e *DefaultDiscoveryEngine) Run(pctx *context.PlatformContext) (*models.DiscoveryManifest, error) {
	logger := pctx.Logger.With("engine", "discovery")
	bus := pctx.EventBus

	logger.Info("Discovery engine starting")
	bus.Publish(events.DiscoveryStarted, nil)

	// Detect platform
	plat, err := detector.NewDetector().Detect()
	if err != nil {
		return nil, err
	}

	// Build the discovery-specific context from PlatformContext.
	dctx := NewContext(logger, bus, pctx.Config, pctx.DB, plat)

	pipeline := NewPipeline(e.config, logger, bus)
	stages := e.registry.All()
	if err := pipeline.AddStages(stages); err != nil {
		return nil, err
	}

	goCtx := pctx.GoContext()
	result, err := pipeline.Run(goCtx, dctx)

	bus.Publish(events.DiscoveryFinished, map[string]string{
		"success":  boolToString(result != nil && result.Success),
		"duration": result.Duration.String(),
	})

	logger.Info(
		"Discovery engine finished",
		"success", result != nil && result.Success,
	)

	// Map Result to DiscoveryManifest
	manifest := &models.DiscoveryManifest{
		ID:        "run-" + plat.ID(), // Basic ID for now
		StartTime: result.StartTime,
		EndTime:   result.EndTime,
		Duration:  result.Duration,
		Platform:  plat.ID(),
		Stages:    make([]models.StageExecutionResult, 0, len(result.Stages)),
		Artifacts: make(map[string]any),
	}

	for name, sr := range result.Stages {
		manifest.Stages = append(manifest.Stages, models.StageExecutionResult{
			Name:     name,
			Status:   string(sr.Status),
			Error:    sr.Error,
			Duration: sr.Duration,
		})
		if sr.Artifact != nil {
			manifest.Artifacts[name] = sr.Artifact
		}
	}

	return manifest, err
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
