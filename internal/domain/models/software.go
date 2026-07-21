package models

// SoftwareInfo represents the overall installed software context.
type SoftwareInfo struct {
	Packages []SoftwarePackage    `json:"packages,omitempty"`
	Runtimes []RuntimeEnvironment `json:"runtimes,omitempty"`
}

// ArtifactType implements discovery.DiscoveryArtifact.
func (SoftwareInfo) ArtifactType() string {
	return "software"
}

// SoftwarePackage represents an installed OS package (e.g. via apt, brew, apk).
type SoftwarePackage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Manager string `json:"manager"` // apt, brew, apk, pacman, etc.
}

// RuntimeEnvironment represents a programming or container runtime.
type RuntimeEnvironment struct {
	Name    string `json:"name"`    // "go", "python", "node", "docker"
	Version string `json:"version"` // "1.21.0"
	Path    string `json:"path"`    // "/usr/local/go/bin/go"
}
