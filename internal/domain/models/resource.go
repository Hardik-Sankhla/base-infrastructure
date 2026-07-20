package models

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
