package hardware

import (
	"context"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// Provider abstracts the platform-specific implementation of hardware discovery.
// This allows for clean mocking in tests and easier porting to new platforms.
type Provider interface {
	GetCPU(ctx context.Context) (models.CPU, error)
	GetRAM(ctx context.Context) (models.RAM, error)
	GetStorage(ctx context.Context) ([]models.Disk, error)
	GetGPUs(ctx context.Context) ([]models.GPU, error)
	GetBattery(ctx context.Context) (models.Battery, error)
	GetThermal(ctx context.Context) ([]models.ThermalSensor, error)
}
