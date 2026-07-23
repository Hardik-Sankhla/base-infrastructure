package discovery

import (
	"context"
	"log/slog"
	"testing"

	"github.com/base-infrastructure/platform/internal/core"
	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform/mock"
)

func TestHardwareStage_Success(t *testing.T) {
	p := mock.NewPlatform()
	p.MockHardware.CPU = models.CPU{Vendor: "GenuineIntel", Architecture: "amd64"}
	p.MockHardware.RAM = models.RAM{TotalBytes: 1024}
	p.MockHardware.Storage = []models.Disk{{Name: "/dev/sda1", Capacity: 500}}

	stage := &HardwareStage{}

	if stage.Name() != "hardware" {
		t.Errorf("expected name 'hardware', got %s", stage.Name())
	}

	ctx, cancel := context.WithTimeout(context.Background(), stage.Timeout())
	defer cancel()

	dctx := core.NewContext(slog.Default(), nil, nil, nil, p)

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

	hw, ok := artifact.(models.Hardware)
	if !ok {
		t.Fatalf("Expected Hardware artifact")
	}

	if hw.CPU.Vendor != "GenuineIntel" {
		t.Errorf("Expected Vendor GenuineIntel, got %s", hw.CPU.Vendor)
	}

	if err := stage.Cleanup(ctx); err != nil {
		t.Fatalf("Cleanup failed: %v", err)
	}
}
