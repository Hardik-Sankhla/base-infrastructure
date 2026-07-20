package models

import "time"

// OSInfo represents the immutable operating system context
type OSInfo struct {
	OperatingSystem     string    `json:"operating_system"`
	Distribution        string    `json:"distribution"`
	DistributionVersion string    `json:"distribution_version"`
	KernelVersion       string    `json:"kernel_version"`
	KernelArchitecture  string    `json:"kernel_architecture"`
	InitSystem          string    `json:"init_system"`
	PackageManager      string    `json:"package_manager"`
	Libc                string    `json:"libc"`
	Shell               string    `json:"shell"`
	Hostname            string    `json:"hostname"`
	Timezone            string    `json:"timezone"`
	Locale              string    `json:"locale"`
	BootTime            time.Time `json:"boot_time"`
}

// ArtifactType implements discovery.DiscoveryArtifact
func (o OSInfo) ArtifactType() string {
	return "OS"
}

// Environment represents the execution context
type Environment struct {
	Virtualization   string `json:"virtualization"`    // wsl, docker, kvm, none
	ContainerRuntime string `json:"container_runtime"` // docker, podman, containerd
}
