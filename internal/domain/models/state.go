package models

// StateManifest represents a parsed desired state configuration file
type StateManifest struct {
	Version      string              `yaml:"version" json:"version"`
	Capabilities []DesiredCapability `yaml:"capabilities" json:"capabilities"`
	Settings     StateSettings       `yaml:"settings,omitempty" json:"settings,omitempty"`
}

// DesiredCapability represents a user's requested capability state
type DesiredCapability struct {
	ID         string            `yaml:"id" json:"id"`
	Provider   string            `yaml:"provider,omitempty" json:"provider,omitempty"`
	Version    string            `yaml:"version,omitempty" json:"version,omitempty"`
	State      CapabilityState   `yaml:"state" json:"state"` // e.g. "available", "missing"
	Properties map[string]string `yaml:"properties,omitempty" json:"properties,omitempty"`
}

// StateSettings holds global settings for how the state should be applied
type StateSettings struct {
	RollbackOnFailure  bool `yaml:"rollback_on_failure" json:"rollback_on_failure"`
	StrictVersionMatch bool `yaml:"strict_version_match" json:"strict_version_match"`
}

// DriftType categorizes the discrepancy between desired and current state
type DriftType string

const (
	DriftMissing  DriftType = "missing"
	DriftVersion  DriftType = "version_mismatch"
	DriftOrphaned DriftType = "orphaned" // Exists but not desired (if enforcing strict state)
)

// Drift represents a required state change
type Drift struct {
	CapabilityID string            `json:"capability_id"`
	Type         DriftType         `json:"type"`
	Desired      DesiredCapability `json:"desired"`
	Current      *Capability       `json:"current,omitempty"` // nil if completely missing
}
