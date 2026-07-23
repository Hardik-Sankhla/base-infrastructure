package hardware

import (
	"github.com/base-infrastructure/platform/internal/runtime"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

// DefaultProvider implements Provider using gopsutil for cross-platform support.
type DefaultProvider struct{}

func NewDefaultProvider() *DefaultProvider {
	return &DefaultProvider{}
}

func (p *DefaultProvider) GetCPU(ctx runtime.Context) (models.CPU, error) {
	var hwCPU models.CPU

	infoStat, err := cpu.Info()
	if err == nil && len(infoStat) > 0 {
		info := infoStat[0]
		hwCPU.Vendor = info.VendorID
		hwCPU.Model = info.ModelName
		hwCPU.PhysicalCores = int(info.Cores)
		hwCPU.CacheL2 = int64(info.CacheSize)
		hwCPU.Flags = info.Flags
	}

	// For threads and logical cores, we use count
	logical, err := cpu.Counts(true)
	if err == nil {
		hwCPU.LogicalCores = logical
		hwCPU.Threads = logical
	}

	physical, err := cpu.Counts(false)
	if err == nil && hwCPU.PhysicalCores == 0 {
		hwCPU.PhysicalCores = physical
	}

	// Basic architecture inference based on Go env or gopsutil fallback
	// This can be enriched with GOARCH in the actual binary
	hwCPU.Architecture = "amd64" // default fallback, ideally populated via runtime.GOARCH

	return hwCPU, nil
}

func (p *DefaultProvider) GetRAM(ctx runtime.Context) (models.RAM, error) {
	v, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return models.RAM{}, err
	}

	s, err := mem.SwapMemoryWithContext(ctx)
	if err != nil {
		return models.RAM{}, err
	}

	return models.RAM{
		TotalBytes:     int64(v.Total),
		AvailableBytes: int64(v.Available),
		UsedBytes:      int64(v.Used),
		SwapTotal:      int64(s.Total),
		SwapUsed:       int64(s.Used),
	}, nil
}

func (p *DefaultProvider) GetStorage(ctx runtime.Context) ([]models.Disk, error) {
	partitions, err := disk.PartitionsWithContext(ctx, true)
	if err != nil {
		return nil, err
	}

	var disks []models.Disk
	for _, p := range partitions {
		usage, err := disk.UsageWithContext(ctx, p.Mountpoint)
		if err != nil {
			continue // skip inaccessible partitions
		}

		disks = append(disks, models.Disk{
			Name:       p.Device,
			Type:       "unknown", // gopsutil doesn't natively expose SSD/HDD easily
			Capacity:   int64(usage.Total),
			FreeSpace:  int64(usage.Free),
			Filesystem: p.Fstype,
			MountPoint: p.Mountpoint,
		})
	}

	return disks, nil
}

func (p *DefaultProvider) GetGPUs(ctx runtime.Context) ([]models.GPU, error) {
	// gopsutil does not have built-in GPU discovery.
	// This would require executing lspci, nvidia-smi, or querying WMI/IOKit.
	// Returning empty list gracefully.
	return nil, nil
}

func (p *DefaultProvider) GetBattery(ctx runtime.Context) (models.Battery, error) {
	// gopsutil does not have reliable battery polling.
	return models.Battery{Present: false}, nil
}

func (p *DefaultProvider) GetThermal(ctx runtime.Context) ([]models.ThermalSensor, error) {
	// gopsutil does not have reliable cross-platform thermal sensors yet.
	return nil, nil
}
