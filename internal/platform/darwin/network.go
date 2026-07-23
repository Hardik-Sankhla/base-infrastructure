package darwin

import (
	"net"
	"os"
	"strings"

	"github.com/base-infrastructure/platform/internal/runtime"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

type NetworkProvider struct{}

func NewNetworkProvider() *NetworkProvider {
	return &NetworkProvider{}
}

func (p *NetworkProvider) GetInterfaces(ctx runtime.Context) ([]models.NetworkInterface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var result []models.NetworkInterface
	for _, i := range ifaces {
		ni := models.NetworkInterface{
			Name:       i.Name,
			MAC:        i.HardwareAddr.String(),
			MTU:        i.MTU,
			IsUp:       i.Flags&net.FlagUp != 0,
			IsLoopback: i.Flags&net.FlagLoopback != 0,
		}

		addrs, err := i.Addrs()
		if err == nil {
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}

				if ip == nil {
					continue
				}
				if ip.To4() != nil {
					ni.IPv4 = append(ni.IPv4, ip.String())
				} else {
					ni.IPv6 = append(ni.IPv6, ip.String())
				}
			}
		}

		result = append(result, ni)
	}

	return result, nil
}

func (p *NetworkProvider) GetDNS(ctx runtime.Context) (models.DNSConfig, error) {
	config := models.DNSConfig{}

	data, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		// Just return empty if we can't read it
		return config, nil
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		if parts[0] == "nameserver" {
			config.Servers = append(config.Servers, parts[1])
		} else if parts[0] == "search" {
			config.SearchDomains = append(config.SearchDomains, parts[1:]...)
		}
	}

	return config, nil
}

func (p *NetworkProvider) GetProxy(ctx runtime.Context) (models.ProxyConfig, error) {
	return models.ProxyConfig{
		HTTPProxy:  os.Getenv("HTTP_PROXY"),
		HTTPSProxy: os.Getenv("HTTPS_PROXY"),
		NoProxy:    os.Getenv("NO_PROXY"),
	}, nil
}
