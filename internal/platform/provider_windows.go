//go:build windows

package platform

import (
	"context"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform/providers/hardware"
	"github.com/shirou/gopsutil/v3/disk"
)

type windowsEnvironmentProvider struct{}

func NewEnvironmentProvider() *windowsEnvironmentProvider {
	return &windowsEnvironmentProvider{}
}

func (p *windowsEnvironmentProvider) GetEnvironmentInfo(ctx context.Context) (models.EnvironmentInfo, error) {
	var info models.EnvironmentInfo

	// Terminal check
	if stat, err := os.Stdout.Stat(); err == nil {
		info.IsTerminal = (stat.Mode() & os.ModeCharDevice) != 0
	}

	// Root/Admin check (approximate without syscalls)
	if currentUser, err := user.Current(); err == nil {
		// A common hack in Go on windows is checking if we have access to open the PhysicalDrive0,
		// but checking a basic heuristic is often enough for simple contexts.
		// We'll leave IsRoot=false by default unless we implement a dedicated syscall.
		_ = currentUser
	}

	// Container check
	// DOTNET_RUNNING_IN_CONTAINER is common for Windows containers
	if os.Getenv("DOTNET_RUNNING_IN_CONTAINER") == "true" {
		info.IsContainer = true
		info.ContainerRuntime = "docker" // Defaulting for Windows
	}

	// CI Check
	if os.Getenv("CI") != "" {
		info.IsCI = true
		if os.Getenv("GITHUB_ACTIONS") == "true" {
			info.CIProvider = "github"
		} else if strings.EqualFold(os.Getenv("GITLAB_CI"), "true") {
			info.CIProvider = "gitlab"
		} else {
			info.CIProvider = "unknown"
		}
	}

	return info, nil
}

type windowsFilesystemProvider struct{}

func NewFilesystemProvider() *windowsFilesystemProvider {
	return &windowsFilesystemProvider{}
}

func (p *windowsFilesystemProvider) GetFilesystemInfo(ctx context.Context) (models.FilesystemInfo, error) {
	var info models.FilesystemInfo

	// 1. Mounts
	parts, err := disk.PartitionsWithContext(ctx, true)
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

	// 2. Capacity (Use C: by default if available)
	rootDrive := "C:"
	if sysDrive := os.Getenv("SystemDrive"); sysDrive != "" {
		rootDrive = sysDrive
	}
	usage, err := disk.UsageWithContext(ctx, rootDrive+"\\")
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

	if appData := os.Getenv("APPDATA"); appData != "" {
		info.ConfigDir = appData
	} else if home != "" {
		info.ConfigDir = filepath.Join(home, "AppData", "Roaming")
	}

	if localAppData := os.Getenv("LOCALAPPDATA"); localAppData != "" {
		info.DataDir = localAppData
	} else if home != "" {
		info.DataDir = filepath.Join(home, "AppData", "Local")
	}

	info.TempDir = os.TempDir()

	if programData := os.Getenv("PROGRAMDATA"); programData != "" {
		info.RuntimeDir = programData
	} else {
		info.RuntimeDir = "C:\\ProgramData"
	}

	// 4. Executable Search Paths
	if path := os.Getenv("PATH"); path != "" {
		info.SearchPaths = strings.Split(path, string(os.PathListSeparator))
	}

	// 5. Traits
	info.SupportsSymlink = true // Technically requires dev mode or admin, but NTFS supports it
	info.CaseSensitive = false

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

type windowsNetworkProvider struct {
	SharedNetworkProvider
}

func NewNetworkProvider() *windowsNetworkProvider {
	return &windowsNetworkProvider{}
}

func (p *windowsNetworkProvider) GetDNS(ctx context.Context) (models.DNSConfig, error) {
	// Fallback implementation for Windows, to be enhanced later.
	// WMI or Registry reads usually required on Windows.
	return models.DNSConfig{}, nil
}

type windowsOSProvider struct {
	SharedOSProvider
}

func NewOSProvider() *windowsOSProvider {
	return &windowsOSProvider{
		SharedOSProvider: SharedOSProvider{
			OperatingSystem: "windows",
			InitSystem:      "wininit",
			PackageManager:  "unknown",
			Libc:            "msvcrt",
			Shell:           "cmd",
		},
	}
}

// SoftwareProvider implements SoftwareProvider for Windows.
type windowsSoftwareProvider struct{}

// NewSoftwareProvider creates a new Windows software provider.
func NewSoftwareProvider() *windowsSoftwareProvider {
	return &windowsSoftwareProvider{}
}

// GetSoftwareInfo retrieves installed packages and runtimes.
func (p *windowsSoftwareProvider) GetSoftwareInfo(ctx context.Context) (models.SoftwareInfo, error) {
	// TODO: implement winget/choco/registry software discovery
	return models.SoftwareInfo{
		Packages: []models.SoftwarePackage{},
		Runtimes: []models.RuntimeEnvironment{},
	}, nil
}

// Platform implements Platform for Windows.
type windowsPlatform struct {
	BasePlatform
	osProvider       OSProvider
	hardwareProvider HardwareProvider
	fsProvider       FilesystemProvider
}

// NewPlatform creates a new Windows platform instance.
func NewPlatform() *windowsPlatform {
	return &windowsPlatform{
		osProvider:       NewOSProvider(),
		hardwareProvider: hardware.NewDefaultProvider(),
		fsProvider:       NewFilesystemProvider(),
	}
}

func (p *windowsPlatform) ID() string {
	return "windows"
}

func (p *windowsPlatform) Name() string {
	return "Windows"
}

func (p *windowsPlatform) OS() OSProvider {
	return p.osProvider
}

func (p *windowsPlatform) Hardware() HardwareProvider {
	return p.hardwareProvider
}

func (p *windowsPlatform) Filesystem() FilesystemProvider {
	return p.fsProvider
}

func (p *windowsPlatform) Network() NetworkProvider {
	return NewNetworkProvider()
}

func (p *windowsPlatform) Environment() EnvironmentProvider {
	return NewEnvironmentProvider()
}

func (p *windowsPlatform) Software() SoftwareProvider {
	return NewSoftwareProvider()
}
