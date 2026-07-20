package hardware

import (
	"context"
	"log/slog"
	"testing"

	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/domain/models"
)

// mockContext implements discovery.Context for testing
type mockContext struct{}

func (c *mockContext) Logger() *slog.Logger  { return slog.Default() }
func (c *mockContext) EventBus() interface{} { return nil } // Use any to satisfy but we might need real eventbus
func (c *mockContext) Config() interface{}   { return nil }
func (c *mockContext) DB() interface{}       { return nil }

func TestHardwareStage_Success(t *testing.T) {
	mockProv := &MockProvider{
		CPU: models.CPU{Vendor: "GenuineIntel", Architecture: "amd64"},
		RAM: models.RAM{TotalBytes: 1024},
		Storage: []models.Disk{
			{Name: "/dev/sda1", Capacity: 500},
		},
	}

	stage := NewStage(mockProv)

	if stage.Name() != "hardware" {
		t.Errorf("expected name 'hardware', got %s", stage.Name())
	}

	ctx, cancel := context.WithTimeout(context.Background(), stage.Timeout())
	defer cancel()

	// 1. Initialize
	// We need a real struct that implements Context, but the mockContext signatures don't exactly match
	// unless we implement the exact interface from discovery.Context.
	// Let's use the real discovery.NewContext for simplicity, passing nil where allowed.
	dctx := discovery.NewContext(slog.Default(), nil, nil, nil, nil)

	if err := stage.Initialize(dctx); err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}

	// 2. Run
	artifact, err := stage.Run(ctx, dctx)
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}

	// 3. Validate
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

	// 4. Cleanup
	if err := stage.Cleanup(ctx); err != nil {
		t.Fatalf("Cleanup failed: %v", err)
	}
}
