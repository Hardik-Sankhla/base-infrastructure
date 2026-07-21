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



// StageExecutionResult summarizes the outcome of a single discovery stage.
type StageExecutionResult struct {
	Name     string        `json:"name"`
	Status   string        `json:"status"` // success, failed, skipped
	Error    string        `json:"error,omitempty"`
	Duration time.Duration `json:"duration"`
}

// DiscoveryManifest is an immutable artifact summarizing a complete discovery run.
type DiscoveryManifest struct {
	ID        string                 `json:"id"`
	StartTime time.Time              `json:"start_time"`
	EndTime   time.Time              `json:"end_time"`
	Duration  time.Duration          `json:"duration"`
	Platform  string                 `json:"platform"`
	Stages    []StageExecutionResult `json:"stages"`
	Artifacts map[string]any         `json:"artifacts"` // The actual discovered data
}

// MountPoint represents a mounted filesystem.
type MountPoint struct {
	Device     string `json:"device"`
	MountPath  string `json:"mount_path"`
	FSType     string `json:"fs_type"`
	Options    string `json:"options"`
	IsReadOnly bool   `json:"is_read_only"`
}

// FilesystemCapacity represents storage capacity in bytes.
type FilesystemCapacity struct {
	TotalBytes uint64 `json:"total_bytes"`
	UsedBytes  uint64 `json:"used_bytes"`
	FreeBytes  uint64 `json:"free_bytes"`
}

// FilesystemInfo contains immutable facts about the system's filesystems and paths.
type FilesystemInfo struct {
	Mounts          []MountPoint       `json:"mounts"`
	RootCapacity    FilesystemCapacity `json:"root_capacity"`
	HomeDir         string             `json:"home_dir"`
	ConfigDir       string             `json:"config_dir"`
	DataDir         string             `json:"data_dir"`
	TempDir         string             `json:"temp_dir"`
	RuntimeDir      string             `json:"runtime_dir"`
	SearchPaths     []string           `json:"search_paths"` // PATH
	SupportsSymlink bool               `json:"supports_symlink"`
	CaseSensitive   bool               `json:"case_sensitive"`
}

// ArtifactType implements discovery.DiscoveryArtifact
func (f FilesystemInfo) ArtifactType() string {
	return "Filesystem"
}
