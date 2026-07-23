package discovery

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/base-infrastructure/platform/internal/core"
	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform/mock"
)

func TestEnvironmentStage_Success(t *testing.T) {
	stage := &EnvironmentStage{}
	p := mock.NewPlatform()

	p.MockEnvironment.Info = models.EnvironmentInfo{
		IsVirtualMachine: true,
		Virtualization:   "wsl",
		IsTerminal:       true,
	}

	log := slog.Default()
	ctx := core.NewContext(log, nil, nil, nil, p)

	if err := stage.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize: %v", err)
	}

	goCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	artifact, err := stage.Run(goCtx, ctx)
	if err != nil {
		t.Fatalf("core.Stage run failed: %v", err)
	}

	if err := stage.Validate(artifact); err != nil {
		t.Fatalf("Artifact validation failed: %v", err)
	}

	envInfo, ok := artifact.(models.EnvironmentInfo)
	if !ok {
		t.Fatalf("Expected models.EnvironmentInfo, got %T", artifact)
	}

	if !envInfo.IsVirtualMachine || envInfo.Virtualization != "wsl" {
		t.Errorf("Unexpected environment info: %+v", envInfo)
	}
}

func TestEnvironmentStage_UninitializedPlatform(t *testing.T) {
	stage := &EnvironmentStage{}
	log := slog.Default()

	// core.Context with NO platform
	ctx := core.NewContext(log, nil, nil, nil, nil)

	err := stage.Initialize(ctx)
	if err == nil {
		t.Fatal("Expected initialization to fail with no platform")
	}
}

func TestEnvironmentStage_GetEnvironmentError(t *testing.T) {
	stage := &EnvironmentStage{}
	p := mock.NewPlatform()

	p.MockEnvironment.Err = errors.New("environment provider error")

	log := slog.Default()
	ctx := core.NewContext(log, nil, nil, nil, p)

	if err := stage.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize: %v", err)
	}

	goCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := stage.Run(goCtx, ctx)
	if err == nil {
		t.Fatal("Expected run to fail when environment error occurs")
	}
}
