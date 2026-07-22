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
			sb.WriteString(fmt.Sprintf("  %s\n", osInfo.KernelArchitecture))
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
				sb.WriteString(fmt.Sprintf("  %d Devices\n", len(hwInfo.Storage)))
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
				sb.WriteString(fmt.Sprintf("  %d Mounts\n", len(fsInfo.Mounts)))
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
	sb.WriteString(fmt.Sprintf("  %d discovered\n", len(res.Capabilities)))
	if verbosity > 0 {
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

func getFloat(v any) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case uint64:
		return float64(val)
	default:
		return 0
	}
}
