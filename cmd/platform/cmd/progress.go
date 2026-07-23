package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/base-infrastructure/platform/internal/discovery"
)

// ProgressHook implements discovery.Hook to print human-readable progress.
type ProgressHook struct {
	startTimes map[string]time.Time
}

func NewProgressHook() *ProgressHook {
	return &ProgressHook{
		startTimes: make(map[string]time.Time),
	}
}

func (h *ProgressHook) BeforePipeline(ctx context.Context, dctx discovery.Context, stages []discovery.Stage) error {
	fmt.Println("Discovering platform...")
	return nil
}

func (h *ProgressHook) AfterPipeline(ctx context.Context, dctx discovery.Context, results map[string]discovery.StageResult) error {
	return nil
}

func (h *ProgressHook) BeforeStage(ctx context.Context, dctx discovery.Context, stage discovery.Stage) error {
	h.startTimes[stage.Name()] = time.Now()
	fmt.Printf("Discovering %s...\n", stage.Name())
	return nil
}

func (h *ProgressHook) AfterStage(ctx context.Context, dctx discovery.Context, stage discovery.Stage, artifact discovery.DiscoveryArtifact) error {
	start := h.startTimes[stage.Name()]
	elapsed := time.Since(start)

	durStr := fmt.Sprintf("%dms", elapsed.Milliseconds())
	if elapsed >= time.Second {
		durStr = fmt.Sprintf("%.2fs", elapsed.Seconds())
	}

	// Capitalize stage name for display
	name := stage.Name()
	if len(name) > 0 {
		name = string(name[0]-32) + name[1:] // simple ASCII title case since names are like 'hardware'
	}

	fmt.Printf("  ✓ %s (%s)\n", name, durStr)
	return nil
}

func (h *ProgressHook) OnStageError(ctx context.Context, dctx discovery.Context, stage discovery.Stage, err error) {
	fmt.Printf("✗ %s failed: %v\n", stage.Name(), err)
}
