package platform

import (
	"context"
	"net"
	"os"
	"strings"
	"time"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/shirou/gopsutil/v3/host"
)

type SharedNetworkProvider struct{}

func (p *SharedNetworkProvider) GetInterfaces(ctx context.Context) ([]models.NetworkInterface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var result []models.NetworkInterface
	for _, i := range ifaces {
		ni := models.NetworkInterface{
			Name:       i.Name,
			MAC:        i.HardwareAddr.String(),
			MTU:        i.MTU,
			IsUp:       i.Flags&net.FlagUp != 0,
			IsLoopback: i.Flags&net.FlagLoopback != 0,
		}

		addrs, err := i.Addrs()
		if err == nil {
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}

				if ip == nil {
					continue
				}
				if ip.To4() != nil {
					ni.IPv4 = append(ni.IPv4, ip.String())
				} else {
					ni.IPv6 = append(ni.IPv6, ip.String())
				}
			}
		}

		result = append(result, ni)
	}

	return result, nil
}

func (p *SharedNetworkProvider) GetProxy(ctx context.Context) (models.ProxyConfig, error) {
	return models.ProxyConfig{
		HTTPProxy:  getEnvIgnoreCase("HTTP_PROXY"),
		HTTPSProxy: getEnvIgnoreCase("HTTPS_PROXY"),
		NoProxy:    getEnvIgnoreCase("NO_PROXY"),
	}, nil
}

func getEnvIgnoreCase(key string) string {
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], key) {
			return parts[1]
		}
	}
	return ""
}

type SharedOSProvider struct {
	OperatingSystem string
	InitSystem      string
	PackageManager  string
	Libc            string
	Shell           string
}

func (p *SharedOSProvider) GetOSInfo(ctx context.Context) (models.OSInfo, error) {
	var info models.OSInfo
	info.OperatingSystem = p.OperatingSystem

	hInfo, err := host.InfoWithContext(ctx)
	if err == nil {
		info.Distribution = hInfo.Platform
		info.DistributionVersion = hInfo.PlatformVersion
		info.KernelVersion = hInfo.KernelVersion
		info.KernelArchitecture = hInfo.KernelArch
		info.Hostname = hInfo.Hostname
		info.BootTime = time.Unix(int64(hInfo.BootTime), 0)
	}

	info.InitSystem = p.InitSystem
	info.PackageManager = p.PackageManager
	info.Libc = p.Libc
	info.Shell = p.Shell

	return info, nil
}
