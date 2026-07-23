package linux

import (
	"context"
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

func (p *FilesystemProvider) GetFilesystemInfo(ctx context.Context) (models.FilesystemInfo, error) {
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
