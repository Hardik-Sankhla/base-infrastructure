package mock

import (
	"github.com/base-infrastructure/platform/internal/runtime"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform"
)

type Platform struct {
	platform.BasePlatform
	MockID          string
	MockName        string
	MockOS          *OSProvider
	MockHardware    *HardwareProvider
	MockFilesystem  *FilesystemProvider
	MockNetwork     *NetworkProvider
	MockEnvironment *EnvironmentProvider
	MockSoftware    *SoftwareProvider
}

func NewPlatform() *Platform {
	return &Platform{
		MockID:          "mock",
		MockName:        "Mock OS",
		MockOS:          &OSProvider{},
		MockHardware:    &HardwareProvider{},
		MockFilesystem:  &FilesystemProvider{},
		MockNetwork:     &NetworkProvider{},
		MockEnvironment: &EnvironmentProvider{},
		MockSoftware:    &SoftwareProvider{},
	}
}

func (p *Platform) ID() string                                { return p.MockID }
func (p *Platform) Name() string                              { return p.MockName }
func (p *Platform) OS() platform.OSProvider                   { return p.MockOS }
func (p *Platform) Hardware() platform.HardwareProvider       { return p.MockHardware }
func (p *Platform) Filesystem() platform.FilesystemProvider   { return p.MockFilesystem }
func (p *Platform) Network() platform.NetworkProvider         { return p.MockNetwork }
func (p *Platform) Environment() platform.EnvironmentProvider { return p.MockEnvironment }
func (p *Platform) Software() platform.SoftwareProvider       { return p.MockSoftware }

type EnvironmentProvider struct {
	Info models.EnvironmentInfo
	Err  error
}

func (p *EnvironmentProvider) GetEnvironmentInfo(ctx runtime.Context) (models.EnvironmentInfo, error) {
	return p.Info, p.Err
}

type NetworkProvider struct {
	Interfaces []models.NetworkInterface
	DNS        models.DNSConfig
	Proxy      models.ProxyConfig
	Err        error
}

func (p *NetworkProvider) GetInterfaces(ctx runtime.Context) ([]models.NetworkInterface, error) {
	return p.Interfaces, p.Err
}

func (p *NetworkProvider) GetDNS(ctx runtime.Context) (models.DNSConfig, error) {
	return p.DNS, p.Err
}

func (p *NetworkProvider) GetProxy(ctx runtime.Context) (models.ProxyConfig, error) {
	return p.Proxy, p.Err
}

type FilesystemProvider struct {
	Info models.FilesystemInfo
	Err  error
}

func (p *FilesystemProvider) GetFilesystemInfo(ctx runtime.Context) (models.FilesystemInfo, error) {
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

func (p *HardwareProvider) GetCPU(ctx runtime.Context) (models.CPU, error) { return p.CPU, p.Err }
func (p *HardwareProvider) GetRAM(ctx runtime.Context) (models.RAM, error) { return p.RAM, p.Err }
func (p *HardwareProvider) GetStorage(ctx runtime.Context) ([]models.Disk, error) {
	return p.Storage, p.Err
}

func (p *HardwareProvider) GetGPUs(ctx runtime.Context) ([]models.GPU, error) { return p.GPUs, p.Err }

func (p *HardwareProvider) GetBattery(ctx runtime.Context) (models.Battery, error) {
	return p.Battery, p.Err
}

func (p *HardwareProvider) GetThermal(ctx runtime.Context) ([]models.ThermalSensor, error) {
	return p.Thermals, p.Err
}

type OSProvider struct {
	Info models.OSInfo
	Err  error
}

func (p *OSProvider) GetOSInfo(ctx runtime.Context) (models.OSInfo, error) {
	return p.Info, p.Err
}
