package linux

import (
	"context"
	"time"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/shirou/gopsutil/v3/host"
)

type OSProvider struct{}

func NewOSProvider() *OSProvider {
	return &OSProvider{}
}

func (p *OSProvider) GetOSInfo(ctx context.Context) (models.OSInfo, error) {
	var info models.OSInfo
	info.OperatingSystem = "linux"

	hInfo, err := host.InfoWithContext(ctx)
	if err == nil {
		info.Distribution = hInfo.Platform
		info.DistributionVersion = hInfo.PlatformVersion
		info.KernelVersion = hInfo.KernelVersion
		info.KernelArchitecture = hInfo.KernelArch
		info.Hostname = hInfo.Hostname
		info.BootTime = time.Unix(int64(hInfo.BootTime), 0)
	}

	// For InitSystem, PackageManager, Libc, Shell, Timezone, Locale we can do deeper probing or heuristics.
	// We will implement simple stubs for now to meet the contract.
	info.InitSystem = "unknown"
	info.PackageManager = "unknown"
	info.Libc = "glibc"
	info.Shell = "/bin/bash"

	return info, nil
}
