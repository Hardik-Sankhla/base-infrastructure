package presentation

import (
	"fmt"
	"strings"

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
	if osArt, ok := m.Artifacts["OS"]; ok {
		// Convert map[string]interface{} to string safely or marshal/unmarshal
		// Since it's passed through JSON, it's typically a map[string]interface{} 
		// if we aren't careful with types. Assuming we can access the underlying type 
		// if we use reflection, or we can just marshal it to bytes and back to OSInfo.
		// Wait, the Artifact map contains `any`. Let's just handle it via fmt.Sprintf or typed assertion.
		sb.WriteString("Operating System\n")
		
		// Helper to safely extract fields from map[string]any or typed structs
		if osInfo, isTyped := osArt.(*models.OSInfo); isTyped {
			sb.WriteString(fmt.Sprintf("  %s %s\n", osInfo.Distribution, osInfo.DistributionVersion))
			sb.WriteString(fmt.Sprintf("  Kernel %s\n", osInfo.KernelVersion))
			sb.WriteString(fmt.Sprintf("  Architecture %s\n", osInfo.KernelArchitecture))
		} else if m, isMap := osArt.(map[string]any); isMap {
			sb.WriteString(fmt.Sprintf("  %v %v\n", m["distribution"], m["distribution_version"]))
			sb.WriteString(fmt.Sprintf("  Kernel %v\n", m["kernel_version"]))
			sb.WriteString(fmt.Sprintf("  Architecture %v\n", m["kernel_architecture"]))
		} else {
			sb.WriteString("  [Raw Data Present]\n")
		}
		sb.WriteString("\n")
	}

	// Hardware Info
	if hwArt, ok := m.Artifacts["Hardware"]; ok {
		sb.WriteString("Hardware\n")
		if hwMap, isMap := hwArt.(map[string]any); isMap {
			if cpu, ok := hwMap["cpu"].(map[string]any); ok {
				sb.WriteString(fmt.Sprintf("  CPU: %v (%v logical cores)\n", cpu["model"], cpu["logical_cores"]))
			}
			if ram, ok := hwMap["ram"].(map[string]any); ok {
				totalBytes := getFloat(ram["total_bytes"])
				availableBytes := getFloat(ram["available_bytes"])
				sb.WriteString(fmt.Sprintf("  Memory: %.1f GB Total, %.1f GB Available\n", totalBytes/(1<<30), availableBytes/(1<<30)))
			}
		} else {
			sb.WriteString("  [Hardware Data Present]\n")
		}
		sb.WriteString("\n")
	}

	// Filesystem Info
	if fsArt, ok := m.Artifacts["Filesystem"]; ok {
		sb.WriteString("Storage\n")
		if fsMap, isMap := fsArt.(map[string]any); isMap {
			if rc, ok := fsMap["root_capacity"].(map[string]any); ok {
				total := getFloat(rc["total_bytes"])
				free := getFloat(rc["free_bytes"])
				sb.WriteString(fmt.Sprintf("  Root: %.1f GB (%.1f GB Free)\n", total/(1<<30), free/(1<<30)))
			}
		} else {
			sb.WriteString("  [Filesystem Data Present]\n")
		}
		sb.WriteString("\n")
	}

	// Network Info
	if netArt, ok := m.Artifacts["Network"]; ok {
		sb.WriteString("Network\n")
		if netMap, isMap := netArt.(map[string]any); isMap {
			if ifaces, ok := netMap["interfaces"].([]any); ok {
				sb.WriteString(fmt.Sprintf("  Interfaces: %d\n", len(ifaces)))
				for _, i := range ifaces {
					if iface, ok := i.(map[string]any); ok {
						if isUp, _ := iface["is_up"].(bool); isUp {
							if isLoop, _ := iface["is_loopback"].(bool); !isLoop {
								if ips, ok := iface["ipv4"].([]any); ok && len(ips) > 0 {
									sb.WriteString(fmt.Sprintf("  Active: %v (IPv4: %v)\n", iface["name"], ips[0]))
									break
								}
							}
						}
					}
				}
			}
		} else {
			sb.WriteString("  [Network Data Present]\n")
		}
		sb.WriteString("\n")
	}

	// Capabilities
	if len(res.Capabilities) > 0 {
		sb.WriteString("Capabilities\n")
		for _, cap := range res.Capabilities {
			sb.WriteString(fmt.Sprintf("  ✓ %s\n", cap.ID))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("Discovery completed successfully.\n")

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
