package linux

import (
	"context"
	"os"
	"strings"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

type EnvironmentProvider struct{}

func NewEnvironmentProvider() *EnvironmentProvider {
	return &EnvironmentProvider{}
}

func (p *EnvironmentProvider) GetEnvironmentInfo(ctx context.Context) (models.EnvironmentInfo, error) {
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
