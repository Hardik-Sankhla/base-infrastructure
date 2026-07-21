package models

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
