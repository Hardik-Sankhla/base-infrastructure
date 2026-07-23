//go:build linux

package platform

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform/providers/hardware"
	"github.com/shirou/gopsutil/v3/disk"
)

type linuxEnvironmentProvider struct{}

func NewEnvironmentProvider() *linuxEnvironmentProvider {
	return &linuxEnvironmentProvider{}
}

func (p *linuxEnvironmentProvider) GetEnvironmentInfo(ctx context.Context) (models.EnvironmentInfo, error) {
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

type linuxFilesystemProvider struct{}

func NewFilesystemProvider() *linuxFilesystemProvider {
	return &linuxFilesystemProvider{}
}

func (p *linuxFilesystemProvider) GetFilesystemInfo(ctx context.Context) (models.FilesystemInfo, error) {
	var info models.FilesystemInfo

	// 1. Mounts
	parts, err := disk.PartitionsWithContext(ctx, true) // true = all partitions
	if err == nil {
		for _, part := range parts {
			info.Mounts = append(info.Mounts, models.MountPoint{
				Device:     part.Device,
				MountPath:  part.Mountpoint,
				FSType:     part.Fstype,
				Options:    strings.Join(part.Opts, ","),
				IsReadOnly: containsIgnoreCase(part.Opts, "ro"),
			})
		}
	}

	// 2. Capacity
	usage, err := disk.UsageWithContext(ctx, "/")
	if err == nil {
		info.RootCapacity = models.FilesystemCapacity{
			TotalBytes: usage.Total,
			UsedBytes:  usage.Used,
			FreeBytes:  usage.Free,
		}
	}

	// 3. Standard Directories
	home, _ := os.UserHomeDir()
	info.HomeDir = home

	if cfg := os.Getenv("XDG_CONFIG_HOME"); cfg != "" {
		info.ConfigDir = cfg
	} else if home != "" {
		info.ConfigDir = filepath.Join(home, ".config")
	}

	if data := os.Getenv("XDG_DATA_HOME"); data != "" {
		info.DataDir = data
	} else if home != "" {
		info.DataDir = filepath.Join(home, ".local", "share")
	}

	info.TempDir = os.TempDir()

	if runtime := os.Getenv("XDG_RUNTIME_DIR"); runtime != "" {
		info.RuntimeDir = runtime
	} else {
		info.RuntimeDir = "/run"
	}

	// 4. Executable Search Paths
	if path := os.Getenv("PATH"); path != "" {
		info.SearchPaths = strings.Split(path, string(os.PathListSeparator))
	}

	// 5. Traits
	info.SupportsSymlink = true
	info.CaseSensitive = true

	return info, nil
}

func containsIgnoreCase(slice []string, val string) bool {
	lowerVal := strings.ToLower(val)
	for _, item := range slice {
		if strings.ToLower(item) == lowerVal {
			return true
		}
	}
	return false
}

// Platform implements Platform for Linux.
type linuxPlatform struct {
	BasePlatform
	osProvider       OSProvider
	hardwareProvider HardwareProvider
	fsProvider       FilesystemProvider
}

// NewPlatform creates a new Linux platform instance.
func NewPlatform() *linuxPlatform {
	return &linuxPlatform{
		osProvider:       NewOSProvider(),
		hardwareProvider: hardware.NewDefaultProvider(),
		fsProvider:       NewFilesystemProvider(),
	}
}

func (p *linuxPlatform) ID() string {
	return "linux"
}

func (p *linuxPlatform) Name() string {
	return "Linux"
}

func (p *linuxPlatform) OS() OSProvider {
	return p.osProvider
}

func (p *linuxPlatform) Hardware() HardwareProvider {
	return p.hardwareProvider
}

func (p *linuxPlatform) Filesystem() FilesystemProvider {
	return p.fsProvider
}

func (p *linuxPlatform) Network() NetworkProvider {
	return NewNetworkProvider()
}

func (p *linuxPlatform) Environment() EnvironmentProvider {
	return NewEnvironmentProvider()
}

func (p *linuxPlatform) Software() SoftwareProvider {
	return NewSoftwareProvider()
}

type linuxNetworkProvider struct {
	SharedNetworkProvider
}

func NewNetworkProvider() *linuxNetworkProvider {
	return &linuxNetworkProvider{}
}

func (p *linuxNetworkProvider) GetDNS(ctx context.Context) (models.DNSConfig, error) {
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

type linuxOSProvider struct {
	SharedOSProvider
}

func NewOSProvider() *linuxOSProvider {
	return &linuxOSProvider{
		SharedOSProvider: SharedOSProvider{
			OperatingSystem: "linux",
			InitSystem:      "unknown",
			PackageManager:  "unknown",
			Libc:            "glibc",
			Shell:           "/bin/bash",
		},
	}
}

// SoftwareProvider implements SoftwareProvider for Linux.
type linuxSoftwareProvider struct{}

// NewSoftwareProvider creates a new Linux software provider.
func NewSoftwareProvider() *linuxSoftwareProvider {
	return &linuxSoftwareProvider{}
}

// GetSoftwareInfo retrieves installed packages and runtimes.
func (p *linuxSoftwareProvider) GetSoftwareInfo(ctx context.Context) (models.SoftwareInfo, error) {
	// TODO: implement apt/yum/apk software discovery
	return models.SoftwareInfo{
		Packages: []models.SoftwarePackage{},
		Runtimes: []models.RuntimeEnvironment{},
	}, nil
}
