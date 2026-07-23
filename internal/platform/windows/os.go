package windows

import (
	"time"

	"github.com/base-infrastructure/platform/internal/runtime"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/shirou/gopsutil/v3/host"
)

type OSProvider struct{}

func NewOSProvider() *OSProvider {
	return &OSProvider{}
}

func (p *OSProvider) GetOSInfo(ctx runtime.Context) (models.OSInfo, error) {
	var info models.OSInfo
	info.OperatingSystem = "windows"

	hInfo, err := host.InfoWithContext(ctx)
	if err == nil {
		info.Distribution = hInfo.Platform
		info.DistributionVersion = hInfo.PlatformVersion
		info.KernelVersion = hInfo.KernelVersion
		info.KernelArchitecture = hInfo.KernelArch
		info.Hostname = hInfo.Hostname
		info.BootTime = time.Unix(int64(hInfo.BootTime), 0)
	}

	info.InitSystem = "wininit"
	info.PackageManager = "unknown"
	info.Libc = "msvcrt"
	info.Shell = "cmd"

	return info, nil
}
