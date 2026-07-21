package models

// EnvironmentInfo contains immutable facts about the system's execution context.
type EnvironmentInfo struct {
	IsContainer      bool   `json:"is_container"`
	ContainerRuntime string `json:"container_runtime,omitempty"` // docker, podman, containerd, lxc

	IsVirtualMachine bool   `json:"is_virtual_machine"`
	Virtualization   string `json:"virtualization,omitempty"` // wsl, kvm, vmware, hyperv, virtualbox

	IsCloud       bool   `json:"is_cloud"`
	CloudProvider string `json:"cloud_provider,omitempty"` // aws, gcp, azure, digitalocean

	IsCI       bool   `json:"is_ci"`
	CIProvider string `json:"ci_provider,omitempty"` // github, gitlab, jenkins, circleci, travis

	IsRoot     bool `json:"is_root"`
	IsTerminal bool `json:"is_terminal"`
}

// ArtifactType implements discovery.DiscoveryArtifact
func (e EnvironmentInfo) ArtifactType() string {
	return "Environment"
}
