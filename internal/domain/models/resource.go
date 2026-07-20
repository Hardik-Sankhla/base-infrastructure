package models

// CPU represents the physical processor characteristics
type CPU struct {
	Architecture string `json:"architecture"`
	Threads      int    `json:"threads"`
	Endianness   string `json:"endianness"`
}

// RAM represents system memory
type RAM struct {
	TotalBytes int64 `json:"total_bytes"`
	SwapBytes  int64 `json:"swap_bytes"`
}

// GPU represents graphics hardware
type GPU struct {
	Vendor string `json:"vendor"`
	Model  string `json:"model"`
	VRAM   int64  `json:"vram"`
}

// Network represents connectivity interfaces
type Network struct {
	Interfaces       []string `json:"interfaces"`
	HasIPv4          bool     `json:"has_ipv4"`
	HasIPv6          bool     `json:"has_ipv6"`
	IsInternetActive bool     `json:"is_internet_active"`
}

// Hardware represents the combined physical resource environment
type Hardware struct {
	CPU     CPU     `json:"cpu"`
	RAM     RAM     `json:"ram"`
	GPU     []GPU   `json:"gpu,omitempty"`
	Network Network `json:"network"`
}
