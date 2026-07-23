//go:build android

package platform

import (
	"context"
	"os"
	"strings"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

type androidPlatform struct {
	BasePlatform
	osProvider OSProvider
}

func NewPlatform() *androidPlatform {
	return &androidPlatform{osProvider: &OSProviderStub{}}
}

func (p *androidPlatform) ID() string     { return "android" }
func (p *androidPlatform) Name() string   { return "Android" }
func (p *androidPlatform) OS() OSProvider { return p.osProvider }

func (p *androidPlatform) Network() NetworkProvider {
	return NewNetworkProvider()
}

func (p *androidPlatform) Environment() EnvironmentProvider {
	return NewEnvironmentProvider()
}

func (p *androidPlatform) Software() SoftwareProvider {
	return NewSoftwareProvider()
}

type OSProviderStub struct{}

func (s *OSProviderStub) GetOSInfo(ctx context.Context) (models.OSInfo, error) {
	return models.OSInfo{OperatingSystem: "android"}, nil
}

type androidEnvironmentProvider struct{}

func NewEnvironmentProvider() *androidEnvironmentProvider {
	return &androidEnvironmentProvider{}
}

func (p *androidEnvironmentProvider) GetEnvironmentInfo(ctx context.Context) (models.EnvironmentInfo, error) {
	var info models.EnvironmentInfo

	// Terminal check
	if stat, err := os.Stdout.Stat(); err == nil {
		info.IsTerminal = (stat.Mode() & os.ModeCharDevice) != 0
	}

	// Root check
	info.IsRoot = os.Geteuid() == 0

	// Container check
	if _, err := os.Stat("/.dockerenv"); err == nil {
		info.IsContainer = true
		info.ContainerRuntime = "docker"
	} else {
		// Fallback to cgroup
		if cgroup, err := os.ReadFile("/proc/1/cgroup"); err == nil {
			cgroupStr := string(cgroup)
			if strings.Contains(cgroupStr, "docker") {
				info.IsContainer = true
				info.ContainerRuntime = "docker"
			} else if strings.Contains(cgroupStr, "lxc") {
				info.IsContainer = true
				info.ContainerRuntime = "lxc"
			}
		}
	}

	// VM & Cloud check
	if product, err := os.ReadFile("/sys/class/dmi/id/product_name"); err == nil {
		prodStr := strings.ToLower(strings.TrimSpace(string(product)))
		if strings.Contains(prodStr, "kvm") {
			info.IsVirtualMachine = true
			info.Virtualization = "kvm"
		} else if strings.Contains(prodStr, "vmware") {
			info.IsVirtualMachine = true
			info.Virtualization = "vmware"
		} else if strings.Contains(prodStr, "virtualbox") {
			info.IsVirtualMachine = true
			info.Virtualization = "virtualbox"
		}
	}

	if sysVendor, err := os.ReadFile("/sys/class/dmi/id/sys_vendor"); err == nil {
		vendorStr := strings.ToLower(strings.TrimSpace(string(sysVendor)))
		if strings.Contains(vendorStr, "amazon") {
			info.IsCloud = true
			info.CloudProvider = "aws"
		} else if strings.Contains(vendorStr, "google") {
			info.IsCloud = true
			info.CloudProvider = "gcp"
		} else if strings.Contains(vendorStr, "microsoft") {
			info.IsCloud = true
			info.CloudProvider = "azure"
		}
	}

	// WSL check
	if version, err := os.ReadFile("/proc/version"); err == nil {
		if strings.Contains(strings.ToLower(string(version)), "microsoft") {
			info.IsVirtualMachine = true
			info.Virtualization = "wsl"
		}
	}

	// CI Check
	if os.Getenv("CI") != "" {
		info.IsCI = true
		if os.Getenv("GITHUB_ACTIONS") == "true" {
			info.CIProvider = "github"
		} else if os.Getenv("GITLAB_CI") == "true" {
			info.CIProvider = "gitlab"
		} else {
			info.CIProvider = "unknown"
		}
	}

	return info, nil
}

type androidNetworkProvider struct {
	SharedNetworkProvider
}

func NewNetworkProvider() *androidNetworkProvider {
	return &androidNetworkProvider{}
}

func (p *androidNetworkProvider) GetDNS(ctx context.Context) (models.DNSConfig, error) {
	config := models.DNSConfig{}

	data, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		// Just return empty if we can't read it
		return config, nil
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		if parts[0] == "nameserver" {
			config.Servers = append(config.Servers, parts[1])
		} else if parts[0] == "search" {
			config.SearchDomains = append(config.SearchDomains, parts[1:]...)
		}
	}

	return config, nil
}

// SoftwareProvider implements SoftwareProvider for Android.
type androidSoftwareProvider struct{}

// NewSoftwareProvider creates a new Android software provider.
func NewSoftwareProvider() *androidSoftwareProvider {
	return &androidSoftwareProvider{}
}

// GetSoftwareInfo retrieves installed packages and runtimes.
func (p *androidSoftwareProvider) GetSoftwareInfo(ctx context.Context) (models.SoftwareInfo, error) {
	// TODO: implement termux/android software discovery
	return models.SoftwareInfo{
		Packages: []models.SoftwarePackage{},
		Runtimes: []models.RuntimeEnvironment{},
	}, nil
}
