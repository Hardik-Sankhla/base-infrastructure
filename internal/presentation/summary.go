package presentation

import (
	"fmt"
	"strings"
	"time"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

func formatSummary(res Result, verbosity int) string {
	var sb strings.Builder
	m := res.Manifest

	sb.WriteString("\nPlatform Discovery Summary\n")
	sb.WriteString("────────────────────────────────────────\n\n")

	if m == nil {
		sb.WriteString("No discovery manifest provided.\n")
		return sb.String()
	}

	// OS Info
	if osArt, ok := m.Artifacts["os"]; ok {
		sb.WriteString("Host\n")
		if osInfo, isTyped := osArt.(models.OSInfo); isTyped {
			sb.WriteString(fmt.Sprintf("  %s %s\n", osInfo.Distribution, osInfo.DistributionVersion))
			sb.WriteString("Architecture\n")
			
			arch := osInfo.KernelArchitecture
			if arch == "x86_64" {
				arch = "AMD64 (x86_64)"
			} else if arch == "aarch64" {
				arch = "ARM64 (aarch64)"
			}
			sb.WriteString(fmt.Sprintf("  %s\n", arch))
		} else {
			sb.WriteString("  [Raw OS Data Present]\n")
		}
	}

	// Hardware Info
	if hwArt, ok := m.Artifacts["hardware"]; ok {
		if hwInfo, isTyped := hwArt.(models.Hardware); isTyped {
			if hwInfo.CPU.Model != "" {
				sb.WriteString("CPU\n")
				sb.WriteString(fmt.Sprintf("  %s\n", hwInfo.CPU.Model))
			}
			if hwInfo.RAM.TotalBytes > 0 {
				sb.WriteString("Memory\n")
				sb.WriteString(fmt.Sprintf("  %.1f GB\n", float64(hwInfo.RAM.TotalBytes)/(1<<30)))
			}
			if len(hwInfo.Storage) > 0 {
				sb.WriteString("Storage\n")
				var phys, loop, ram, virtual int
				for _, d := range hwInfo.Storage {
					if strings.HasPrefix(d.Name, "/dev/loop") {
						loop++
					} else if strings.HasPrefix(d.Name, "/dev/ram") {
						ram++
					} else if strings.Contains(d.Name, "vd") || strings.Contains(d.Name, "virtual") {
						virtual++
					} else {
						phys++
					}
				}
				if phys > 0 {
					sb.WriteString(fmt.Sprintf("  %d Physical Devices\n", phys))
				}
				if virtual > 0 {
					sb.WriteString(fmt.Sprintf("  %d Virtual Devices\n", virtual))
				}
				if loop > 0 {
					sb.WriteString(fmt.Sprintf("  %d Loop Devices\n", loop))
				}
				if ram > 0 {
					sb.WriteString(fmt.Sprintf("  %d RAM Disks\n", ram))
				}
			}
		} else {
			sb.WriteString("Hardware\n  [Raw Hardware Data Present]\n")
		}
	}

	// Filesystem Info
	if fsArt, ok := m.Artifacts["filesystem"]; ok {
		if fsInfo, isTyped := fsArt.(models.FilesystemInfo); isTyped {
			if len(fsInfo.Mounts) > 0 {
				sb.WriteString("Filesystem\n")
				var real, pseudo int
				for _, m := range fsInfo.Mounts {
					if (strings.HasPrefix(m.Device, "/dev/") && !strings.HasPrefix(m.Device, "/dev/loop") && !strings.HasPrefix(m.Device, "/dev/ram")) || m.FSType == "zfs" || m.FSType == "btrfs" || m.FSType == "xfs" || m.FSType == "ext4" {
						real++
					} else {
						pseudo++
					}
				}
				if real > 0 {
					sb.WriteString(fmt.Sprintf("  %d Real Filesystems\n", real))
				}
				if pseudo > 0 {
					sb.WriteString(fmt.Sprintf("  %d Pseudo Mounts\n", pseudo))
				}
			}
		} else {
			sb.WriteString("Filesystem\n  [Raw Filesystem Data Present]\n")
		}
	}

	// Network Info
	if netArt, ok := m.Artifacts["network"]; ok {
		if netInfo, isTyped := netArt.(models.NetworkInfo); isTyped {
			if len(netInfo.Interfaces) > 0 {
				sb.WriteString("Network\n")
				sb.WriteString(fmt.Sprintf("  %d Interfaces\n", len(netInfo.Interfaces)))
			}
		} else {
			sb.WriteString("Network\n  [Raw Network Data Present]\n")
		}
	}

	sb.WriteString("\nCapabilities\n")
	if len(res.Capabilities) == 0 {
		sb.WriteString("  0 discovered\n")
	} else {
		for _, cap := range res.Capabilities {
			sb.WriteString(fmt.Sprintf("  ✓ %s\n", cap.ID))
		}
	}

	sb.WriteString("\nDiscovery Time\n")
	// Custom format for durations: if >1s, use s, else ms.
	dur := m.Duration
	if dur >= time.Second {
		sb.WriteString(fmt.Sprintf("  %.2f s\n", dur.Seconds()))
	} else {
		sb.WriteString(fmt.Sprintf("  %d ms\n", dur.Milliseconds()))
	}

	return sb.String()
}
