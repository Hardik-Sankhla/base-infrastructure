package presentation

import (
	"strings"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// filterManifest returns a new DiscoveryManifest with only the requested artifacts,
// or the original manifest if no filters are provided.
func filterManifest(m *models.DiscoveryManifest, filters []string) *models.DiscoveryManifest {
	if m == nil || len(filters) == 0 {
		return m
	}

	// Create a fast lookup map for filters (case insensitive)
	filterMap := make(map[string]bool)
	for _, f := range filters {
		filterMap[strings.ToLower(f)] = true
	}

	filteredArtifacts := make(map[string]any)
	for key, val := range m.Artifacts {
		if filterMap[strings.ToLower(key)] {
			filteredArtifacts[key] = val
		}
	}

	// Create a shallow copy of the manifest and replace artifacts
	filtered := *m
	filtered.Artifacts = filteredArtifacts
	return &filtered
}
