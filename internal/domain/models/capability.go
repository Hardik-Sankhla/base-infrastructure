package models

// CapabilityState represents the current state of a capability
type CapabilityState string

const (
	StateAvailable CapabilityState = "available"
	StateMissing   CapabilityState = "missing"
	StateBroken    CapabilityState = "broken"
)

// Capability represents a functional ability rather than a specific software product.
// e.g. ID="container.runtime", Provider="docker"
type Capability struct {
	ID         string            `json:"id"`
	Provider   string            `json:"provider"`
	Version    string            `json:"version"`
	State      CapabilityState   `json:"state"`
	Confidence int               `json:"confidence"` // 0-100
	Metadata   map[string]string `json:"metadata"`
}
