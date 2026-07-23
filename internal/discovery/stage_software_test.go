package discovery

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform/mock"
	"github.com/base-infrastructure/platform/internal/runtime"
)

func TestSoftwareStage_Success(t *testing.T) {
	stage := &SoftwareStage{}
	p := mock.NewPlatform()

	p.MockSoftware.Info = models.SoftwareInfo{
		Packages: []models.SoftwarePackage{
			{Name: "git", Version: "2.34.1", Manager: "apt"},
		},
		Runtimes: []models.RuntimeEnvironment{
			{Name: "go", Version: "1.21.0", Path: "/usr/local/go/bin/go"},
		},
	}

	log := slog.Default()
	bus := runtime.NewEventBus()
	ctx := NewContext(log, bus, nil, nil, p)

	if err := stage.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize: %v", err)
	}

	goCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	artifact, err := stage.Run(goCtx, ctx)
	if err != nil {
		t.Fatalf("Stage run failed: %v", err)
	}

	if err := stage.Validate(artifact); err != nil {
		t.Fatalf("Artifact validation failed: %v", err)
	}

	swInfo, ok := artifact.(models.SoftwareInfo)
	if !ok {
		t.Fatalf("Expected models.SoftwareInfo, got %T", artifact)
	}

	if len(swInfo.Packages) != 1 || swInfo.Packages[0].Name != "git" {
		t.Errorf("Unexpected packages: %+v", swInfo.Packages)
	}
	if len(swInfo.Runtimes) != 1 || swInfo.Runtimes[0].Name != "go" {
		t.Errorf("Unexpected runtimes: %+v", swInfo.Runtimes)
	}
}

func TestSoftwareStage_UninitializedPlatform(t *testing.T) {
	stage := &SoftwareStage{}
	log := slog.Default()
	bus := runtime.NewEventBus()

	ctx := NewContext(log, bus, nil, nil, nil)

	err := stage.Initialize(ctx)
	if err == nil {
		t.Fatal("Expected initialization to fail with no platform")
	}
}

func TestSoftwareStage_ProviderError(t *testing.T) {
	stage := &SoftwareStage{}
	p := mock.NewPlatform()

	p.MockSoftware.Err = errors.New("software provider error")

	log := slog.Default()
	bus := runtime.NewEventBus()
	ctx := NewContext(log, bus, nil, nil, p)

	if err := stage.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize: %v", err)
	}

	goCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := stage.Run(goCtx, ctx)
	if err == nil {
		t.Fatal("Expected run to fail when provider error occurs")
	}
}
