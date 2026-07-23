package windows

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
	// Fallback implementation for Windows, to be enhanced later.
	// WMI or Registry reads usually required on Windows.
	return models.DNSConfig{}, nil
}

func (p *NetworkProvider) GetProxy(ctx runtime.Context) (models.ProxyConfig, error) {
	return models.ProxyConfig{
		HTTPProxy:  getEnvIgnoreCase("HTTP_PROXY"),
		HTTPSProxy: getEnvIgnoreCase("HTTPS_PROXY"),
		NoProxy:    getEnvIgnoreCase("NO_PROXY"),
	}, nil
}

func getEnvIgnoreCase(key string) string {
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], key) {
			return parts[1]
		}
	}
	return ""
}
