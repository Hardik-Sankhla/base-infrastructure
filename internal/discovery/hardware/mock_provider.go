package hardware

import (
	"context"
	"fmt"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// MockProvider provides deterministic hardware data for testing.
type MockProvider struct {
	CPU      models.CPU
	RAM      models.RAM
	Storage  []models.Disk
	GPUs     []models.GPU
	Battery  models.Battery
	Thermals []models.ThermalSensor

	// Allow injecting errors
	ErrCPU     error
	ErrRAM     error
	ErrStorage error
	ErrGPUs    error
	ErrBattery error
	ErrThermal error
}

func (m *MockProvider) GetCPU(ctx context.Context) (models.CPU, error) {
	if m.ErrCPU != nil {
		return models.CPU{}, m.ErrCPU
	}
	return m.CPU, nil
}

func (m *MockProvider) GetRAM(ctx context.Context) (models.RAM, error) {
	if m.ErrRAM != nil {
		return models.RAM{}, m.ErrRAM
	}
	return m.RAM, nil
}

func (m *MockProvider) GetStorage(ctx context.Context) ([]models.Disk, error) {
	if m.ErrStorage != nil {
		return nil, m.ErrStorage
	}
	return m.Storage, nil
}

func (m *MockProvider) GetGPUs(ctx context.Context) ([]models.GPU, error) {
	if m.ErrGPUs != nil {
		return nil, m.ErrGPUs
	}
	// graceful fallback for missing GPU
	if len(m.GPUs) == 0 {
		return nil, fmt.Errorf("no GPU found")
	}
	return m.GPUs, nil
}

func (m *MockProvider) GetBattery(ctx context.Context) (models.Battery, error) {
	if m.ErrBattery != nil {
		return models.Battery{}, m.ErrBattery
	}
	// graceful fallback for missing battery
	if !m.Battery.Present {
		return models.Battery{}, fmt.Errorf("no battery found")
	}
	return m.Battery, nil
}

func (m *MockProvider) GetThermal(ctx context.Context) ([]models.ThermalSensor, error) {
	if m.ErrThermal != nil {
		return nil, m.ErrThermal
	}
	// graceful fallback for missing thermal
	if len(m.Thermals) == 0 {
		return nil, fmt.Errorf("no thermal sensors found")
	}
	return m.Thermals, nil
}
