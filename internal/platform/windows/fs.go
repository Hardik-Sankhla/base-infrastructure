package windows

import (
	"github.com/base-infrastructure/platform/internal/runtime"

	"os"
	"path/filepath"
	"strings"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/shirou/gopsutil/v3/disk"
)

type FilesystemProvider struct{}

func NewFilesystemProvider() *FilesystemProvider {
	return &FilesystemProvider{}
}

func (p *FilesystemProvider) GetFilesystemInfo(ctx runtime.Context) (models.FilesystemInfo, error) {
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
