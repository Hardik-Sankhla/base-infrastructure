package discovery

import (
	"context"
	"fmt"
	"time"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

// Stage implements discovery.Stage for network discovery.
type NetworkStage struct{}

func (s *NetworkStage) Name() string {
	return "network"
}

func (s *NetworkStage) Version() string {
	return "1.0.0"
}

func (s *NetworkStage) Description() string {
	return "Discovers network interfaces, IP addressing, DNS, and proxy configurations"
}

func (s *NetworkStage) Priority() int {
	return 30 // Runs after OS (20)
}

func (s *NetworkStage) DependsOn() []string {
	return []string{"os"} // Soft dependency logically, though can run without it
}

func (s *NetworkStage) Timeout() time.Duration {
	return 30 * time.Second
}

func (s *NetworkStage) Initialize(dctx Context) error {
	if dctx.Platform() == nil {
		return fmt.Errorf("platform abstraction layer is not initialized in context")
	}
	if dctx.Platform().Network() == nil {
		return fmt.Errorf("network provider is not available for this platform")
	}
	return nil
}

func (s *NetworkStage) Run(ctx context.Context, dctx Context) (DiscoveryArtifact, error) {
	var netInfo models.NetworkInfo
	var err error

	provider := dctx.Platform().Network()

	// Retrieve Interfaces
	if netInfo.Interfaces, err = provider.GetInterfaces(ctx); err != nil {
		return nil, fmt.Errorf("failed to discover network interfaces: %w", err)
	}

	// Non-critical: DNS and Proxy can fail gracefully
	if dns, dErr := provider.GetDNS(ctx); dErr == nil {
		netInfo.DNS = dns
	} else {
		dctx.Logger().Debug("Failed to discover DNS configuration", "error", dErr)
	}

	if proxy, pErr := provider.GetProxy(ctx); pErr == nil {
		netInfo.Proxy = proxy
	} else {
		dctx.Logger().Debug("Failed to discover proxy configuration", "error", pErr)
	}

	return netInfo, nil
}

func (s *NetworkStage) Validate(artifact DiscoveryArtifact) error {
	netInfo, ok := artifact.(models.NetworkInfo)
	if !ok {
		return fmt.Errorf("expected models.NetworkInfo artifact, got %T", artifact)
	}

	if len(netInfo.Interfaces) == 0 {
		return fmt.Errorf("invalid artifact: missing network interfaces")
	}

	return nil
}

func (s *NetworkStage) Cleanup(ctx context.Context) error {
	// Nothing to clean up for network discovery
	return nil
}
