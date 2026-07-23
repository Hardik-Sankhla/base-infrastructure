package models

import (
	"time"
)

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

// CompatibilityConstraints dictates version matching logic
type CompatibilityConstraints struct {
	CoreVersion       string `json:"core_version"`
	SDKVersion        string `json:"sdk_version"`
	CapabilityVersion string `json:"capability_version"`
}

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

// NetworkInterface represents a physical or virtual network interface.
type NetworkInterface struct {
	Name       string   `json:"name"`
	MAC        string   `json:"mac"`
	MTU        int      `json:"mtu"`
	IPv4       []string `json:"ipv4"`
	IPv6       []string `json:"ipv6"`
	IsUp       bool     `json:"is_up"`
	IsLoopback bool     `json:"is_loopback"`
}

// DNSConfig represents the system's DNS resolver configuration.
type DNSConfig struct {
	Servers       []string `json:"servers"`
	SearchDomains []string `json:"search_domains"`
}

// ProxyConfig represents the system's configured network proxies.
type ProxyConfig struct {
	HTTPProxy  string `json:"http_proxy"`
	HTTPSProxy string `json:"https_proxy"`
	NoProxy    string `json:"no_proxy"`
}

// NetworkInfo contains immutable facts about the system's network configuration.
type NetworkInfo struct {
	Interfaces []NetworkInterface `json:"interfaces"`
	DNS        DNSConfig          `json:"dns"`
	Proxy      ProxyConfig        `json:"proxy"`
}

// ArtifactType implements discovery.DiscoveryArtifact
func (n NetworkInfo) ArtifactType() string {
	return "Network"
}

// ExecutionPlan represents a strictly versioned execution plan generated by the Planner
type ExecutionPlan struct {
	Version string     `yaml:"version" json:"version"`
	Tasks   []PlanTask `yaml:"tasks" json:"tasks"`
}

// PlanTask represents a single unit of work in the ExecutionPlan
type PlanTask struct {
	ID        string   `yaml:"id" json:"id"`
	Plugin    string   `yaml:"plugin" json:"plugin"`
	DependsOn []string `yaml:"depends_on,omitempty" json:"depends_on,omitempty"`
}

// Policy represents the system configuration policies that influence the planner
type Policy struct {
	AllowBeta          bool   `json:"allow_beta"`
	AllowRoot          bool   `json:"allow_root"`
	InternetRequired   bool   `json:"internet_required"`
	PreferredContainer string `json:"preferred_container"`
	MaxParallelTasks   int    `json:"max_parallel_tasks"`
	OfflineMode        bool   `json:"offline_mode"`
}

// CPU represents the physical processor characteristics
type CPU struct {
	Vendor        string   `json:"vendor"`
	Model         string   `json:"model"`
	Architecture  string   `json:"architecture"`
	PhysicalCores int      `json:"physical_cores"`
	LogicalCores  int      `json:"logical_cores"`
	Threads       int      `json:"threads"`
	CacheL1       int64    `json:"cache_l1,omitempty"`
	CacheL2       int64    `json:"cache_l2,omitempty"`
	CacheL3       int64    `json:"cache_l3,omitempty"`
	Flags         []string `json:"flags,omitempty"`
}

// RAM represents system memory
type RAM struct {
	TotalBytes     int64 `json:"total_bytes"`
	AvailableBytes int64 `json:"available_bytes"`
	UsedBytes      int64 `json:"used_bytes"`
	SwapTotal      int64 `json:"swap_total"`
	SwapUsed       int64 `json:"swap_used"`
}

// Disk represents a storage device
type Disk struct {
	Name       string `json:"name"`
	Type       string `json:"type"` // e.g. HDD, SSD, NVMe
	Capacity   int64  `json:"capacity"`
	FreeSpace  int64  `json:"free_space"`
	Filesystem string `json:"filesystem,omitempty"`
	MountPoint string `json:"mount_point,omitempty"`
}

// GPU represents graphics hardware
type GPU struct {
	Vendor        string `json:"vendor"`
	Model         string `json:"model"`
	DriverVersion string `json:"driver_version,omitempty"`
	VRAM          int64  `json:"vram,omitempty"`
}

// Battery represents a power source
type Battery struct {
	Present       bool    `json:"present"`
	Health        string  `json:"health,omitempty"`         // Good, Fair, Poor
	Capacity      float64 `json:"capacity,omitempty"`       // Percentage
	ChargingState string  `json:"charging_state,omitempty"` // Charging, Discharging, Full
}

// ThermalSensor represents a temperature reading
type ThermalSensor struct {
	Name        string  `json:"name"`
	Temperature float64 `json:"temperature"` // Celsius
}

// Hardware represents the combined physical resource environment
type Hardware struct {
	CPU      CPU             `json:"cpu"`
	RAM      RAM             `json:"ram"`
	Storage  []Disk          `json:"storage"`
	GPUs     []GPU           `json:"gpus,omitempty"`
	Battery  Battery         `json:"battery,omitempty"`
	Thermals []ThermalSensor `json:"thermals,omitempty"`
}

// ArtifactType implements discovery.DiscoveryArtifact
func (h Hardware) ArtifactType() string {
	return "Hardware"
}

// Result is the unified response from all execution engines
type Result struct {
	Success   bool              `json:"success"`
	Warnings  []error           `json:"warnings"`
	Errors    []error           `json:"errors"`
	Duration  time.Duration     `json:"duration"`
	Rollback  bool              `json:"rollback"`
	Artifacts map[string]string `json:"artifacts"`
	Metadata  map[string]string `json:"metadata"`
}

// ManifestSignature represents a cryptographic signature for a plugin manifest
type ManifestSignature struct {
	Algorithm string `json:"algorithm"`
	Hash      string `json:"hash"`
	Signature string `json:"signature"`
	Signer    string `json:"signer"`
}

// TrustedPublisher represents an entity allowed to sign plugins
type TrustedPublisher struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
}

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