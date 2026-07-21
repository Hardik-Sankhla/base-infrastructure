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

	caps = append(caps, b.evaluateNetwork()...)
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
