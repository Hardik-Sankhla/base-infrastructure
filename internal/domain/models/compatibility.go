package models

// CompatibilityConstraints dictates version matching logic
type CompatibilityConstraints struct {
	CoreVersion       string `json:"core_version"`
	SDKVersion        string `json:"sdk_version"`
	CapabilityVersion string `json:"capability_version"`
}
