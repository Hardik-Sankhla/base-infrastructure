package mock

import (
	"context"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform"
)

type Platform struct {
	platform.BasePlatform
	MockID         string
	MockName       string
	MockOS         *OSProvider
	MockHardware   *HardwareProvider
	MockFilesystem *FilesystemProvider
}

func NewPlatform() *Platform {
	return &Platform{
		MockID:         "mock",
		MockName:       "Mock OS",
		MockOS:         &OSProvider{},
		MockHardware:   &HardwareProvider{},
		MockFilesystem: &FilesystemProvider{},
	}
}

func (p *Platform) ID() string              { return p.MockID }
func (p *Platform) Name() string            { return p.MockName }
func (p *Platform) OS() platform.OSProvider { return p.MockOS }
func (p *Platform) Hardware() platform.HardwareProvider { return p.MockHardware }
func (p *Platform) Filesystem() platform.FilesystemProvider { return p.MockFilesystem }

type FilesystemProvider struct {
	Info models.FilesystemInfo
	Err  error
}

func (p *FilesystemProvider) GetFilesystemInfo(ctx context.Context) (models.FilesystemInfo, error) {
	return p.Info, p.Err
}

type HardwareProvider struct {
	CPU      models.CPU
	RAM      models.RAM
	Storage  []models.Disk
	GPUs     []models.GPU
	Battery  models.Battery
	Thermals []models.ThermalSensor
	Err      error
}

func (p *HardwareProvider) GetCPU(ctx context.Context) (models.CPU, error) { return p.CPU, p.Err }
func (p *HardwareProvider) GetRAM(ctx context.Context) (models.RAM, error) { return p.RAM, p.Err }
func (p *HardwareProvider) GetStorage(ctx context.Context) ([]models.Disk, error) { return p.Storage, p.Err }
func (p *HardwareProvider) GetGPUs(ctx context.Context) ([]models.GPU, error) { return p.GPUs, p.Err }
func (p *HardwareProvider) GetBattery(ctx context.Context) (models.Battery, error) { return p.Battery, p.Err }
func (p *HardwareProvider) GetThermal(ctx context.Context) ([]models.ThermalSensor, error) { return p.Thermals, p.Err }

type OSProvider struct {
	Info models.OSInfo
	Err  error
}

func (p *OSProvider) GetOSInfo(ctx context.Context) (models.OSInfo, error) {
	return p.Info, p.Err
}
