package runtime

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type PluginManifest struct {
	SchemaVersion string            `yaml:"schema_version"`
	Name          string            `yaml:"name"`
	Description   string            `yaml:"description"`
	Version       string            `yaml:"version"`
	Compatibility Compatibility     `yaml:"compatibility"`
	Dependencies  []Dependency      `yaml:"dependencies"`
	Provides      []string          `yaml:"provides"`
	Checksums     map[string]string `yaml:"checksums"`
}

type Compatibility struct {
	OS   []string `yaml:"os"`
	Arch []string `yaml:"arch"`
}

type Dependency struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

// LoadManifest reads a manifest.yaml file from disk
func LoadManifest(path string) (*PluginManifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest file: %w", err)
	}

	var m PluginManifest
	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to parse manifest yaml: %w", err)
	}

	// Validate schema
	if m.SchemaVersion == "" {
		return nil, fmt.Errorf("missing schema_version")
	}
	if m.Name == "" {
		return nil, fmt.Errorf("missing plugin name")
	}
	if m.Version == "" {
		return nil, fmt.Errorf("missing plugin version")
	}

	return &m, nil
}
