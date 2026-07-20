package models

// OSInfo represents the operating system context
type OSInfo struct {
	Distribution   string `json:"distribution"` // e.g. Ubuntu, Windows, Debian
	Version        string `json:"version"`
	Kernel         string `json:"kernel"`
	InitSystem     string `json:"init_system"`     // systemd, sysvinit
	PackageManager string `json:"package_manager"` // apt, winget
	Libc           string `json:"libc"`
}

// Environment represents the execution context
type Environment struct {
	Virtualization   string `json:"virtualization"`    // wsl, docker, kvm, none
	ContainerRuntime string `json:"container_runtime"` // docker, podman, containerd
}
