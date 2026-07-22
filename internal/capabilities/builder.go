package capabilities

import (
	"strings"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// Builder translates a DiscoveryManifest into a set of functional Capabilities.
type Builder struct {
	manifest *models.DiscoveryManifest
}

// NewBuilder creates a new capability builder.
func NewBuilder(manifest *models.DiscoveryManifest) *Builder {
	return &Builder{
		manifest: manifest,
	}
}

// Build evaluates the discovery manifest and generates a list of capabilities.
func (b *Builder) Build() []models.Capability {
	var caps []models.Capability

	caps = append(caps, b.evaluateHardware()...)
	caps = append(caps, b.evaluateOS()...)
	caps = append(caps, b.evaluateFilesystem()...)
	caps = append(caps, b.evaluateNetwork()...)
	caps = append(caps, b.evaluateEnvironment()...)
	caps = append(caps, b.evaluateSoftware()...)

	return caps
}

func (b *Builder) evaluateNetwork() []models.Capability {
	var caps []models.Capability
	if netArtifact, ok := b.manifest.Artifacts["network"]; ok {
		if netInfo, ok := netArtifact.(models.NetworkInfo); ok {
			for _, iface := range netInfo.Interfaces {
				if iface.IsUp && len(iface.IPv4) > 0 {
					caps = append(caps, models.Capability{
						ID:         "network.connectivity",
						Provider:   "system",
						State:      models.StateAvailable,
						Confidence: 100,
					})
					break // only need to register it once
				}
			}
		}
	}
	return caps
}

func (b *Builder) evaluateSoftware() []models.Capability {
	var caps []models.Capability
	if swArtifact, ok := b.manifest.Artifacts["software"]; ok {
		if swInfo, ok := swArtifact.(models.SoftwareInfo); ok {
			for _, rt := range swInfo.Runtimes {
				caps = append(caps, models.Capability{
					ID:         "runtime." + strings.ToLower(rt.Name),
					Provider:   rt.Name,
					Version:    rt.Version,
					State:      models.StateAvailable,
					Confidence: 100,
					Metadata:   map[string]string{"path": rt.Path},
				})
				if strings.ToLower(rt.Name) == "docker" {
					caps = append(caps, models.Capability{
						ID:         "container.runtime",
						Provider:   "docker",
						Version:    rt.Version,
						State:      models.StateAvailable,
						Confidence: 100,
					})
				}
			}
		}
	}
	return caps
}

func (b *Builder) evaluateHardware() []models.Capability {
	var caps []models.Capability
	if hwArtifact, ok := b.manifest.Artifacts["hardware"]; ok {
		if hwInfo, ok := hwArtifact.(models.Hardware); ok {
			if hwInfo.CPU.Model != "" {
				caps = append(caps, models.Capability{
					ID: "hardware.cpu", Provider: "system", State: models.StateAvailable, Confidence: 100,
				})
			}
			if hwInfo.RAM.TotalBytes > 0 {
				caps = append(caps, models.Capability{
					ID: "hardware.memory", Provider: "system", State: models.StateAvailable, Confidence: 100,
				})
			}
		}
	}
	return caps
}

func (b *Builder) evaluateOS() []models.Capability {
	var caps []models.Capability
	if osArtifact, ok := b.manifest.Artifacts["os"]; ok {
		if osInfo, ok := osArtifact.(models.OSInfo); ok {
			caps = append(caps, models.Capability{
				ID: "operating-system." + osInfo.OperatingSystem, Provider: "system", State: models.StateAvailable, Confidence: 100,
			})
		}
	}
	return caps
}

func (b *Builder) evaluateFilesystem() []models.Capability {
	var caps []models.Capability
	if fsArtifact, ok := b.manifest.Artifacts["filesystem"]; ok {
		if fsInfo, ok := fsArtifact.(models.FilesystemInfo); ok {
			if len(fsInfo.Mounts) > 0 {
				caps = append(caps, models.Capability{
					ID: "filesystem.local", Provider: "system", State: models.StateAvailable, Confidence: 100,
				})
			}
		}
	}
	return caps
}

func (b *Builder) evaluateEnvironment() []models.Capability {
	var caps []models.Capability
	if envArtifact, ok := b.manifest.Artifacts["environment"]; ok {
		if envInfo, ok := envArtifact.(models.EnvironmentInfo); ok {
			if envInfo.IsVirtualMachine && envInfo.Virtualization != "" {
				caps = append(caps, models.Capability{
					ID: "virtualization." + envInfo.Virtualization, Provider: "system", State: models.StateAvailable, Confidence: 100,
				})
			}
			if envInfo.IsContainer {
				caps = append(caps, models.Capability{
					ID: "environment.container", Provider: "system", State: models.StateAvailable, Confidence: 100,
				})
			}
			if envInfo.IsCI {
				caps = append(caps, models.Capability{
					ID: "environment.ci", Provider: "system", State: models.StateAvailable, Confidence: 100,
				})
			}
		}
	}
	return caps
}
