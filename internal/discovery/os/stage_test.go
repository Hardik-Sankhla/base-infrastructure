package os

import (
	"context"
	"log/slog"
	"testing"

	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform/mock"
)

func TestOSStage_Success(t *testing.T) {
	mockPlat := mock.NewPlatform()
	mockPlat.MockOS.Info = models.OSInfo{
		OperatingSystem: "mock",
		Distribution:    "test-distro",
	}

	stage := NewStage()

	if stage.Name() != "os" {
		t.Errorf("expected name 'os', got %s", stage.Name())
	}

	ctx, cancel := context.WithTimeout(context.Background(), stage.Timeout())
	defer cancel()

	dctx := discovery.NewContext(slog.Default(), nil, nil, nil, mockPlat)

	if err := stage.Initialize(dctx); err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}

	artifact, err := stage.Run(ctx, dctx)
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}

	if err := stage.Validate(artifact); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}

	info, ok := artifact.(models.OSInfo)
	if !ok {
		t.Fatalf("Expected OSInfo artifact")
	}

	if info.OperatingSystem != "mock" {
		t.Errorf("Expected OS mock, got %s", info.OperatingSystem)
	}

	if err := stage.Cleanup(ctx); err != nil {
		t.Fatalf("Cleanup failed: %v", err)
	}
}

func TestOSStage_UninitializedPlatform(t *testing.T) {
	stage := NewStage()
	dctx := discovery.NewContext(slog.Default(), nil, nil, nil, nil)

	err := stage.Initialize(dctx)
	if err == nil {
		t.Fatalf("Expected initialization to fail with nil platform")
	}
}
