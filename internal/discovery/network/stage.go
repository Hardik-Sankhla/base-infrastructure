package network

import (
	"context"
	"fmt"
	"time"

	"github.com/base-infrastructure/platform/internal/discovery"
	"github.com/base-infrastructure/platform/internal/domain/models"
)

// Stage implements discovery.Stage for network discovery.
type Stage struct{}

// NewStage creates a new Network discovery stage.
func NewStage() *Stage {
	return &Stage{}
}

func (s *Stage) Name() string {
	return "network"
}

func (s *Stage) Version() string {
	return "1.0.0"
}

func (s *Stage) Description() string {
	return "Discovers network interfaces, IP addressing, DNS, and proxy configurations"
}

func (s *Stage) Priority() int {
	return 30 // Runs after OS (20)
}

func (s *Stage) DependsOn() []string {
	return []string{"os"} // Soft dependency logically, though can run without it
}

func (s *Stage) Timeout() time.Duration {
	return 30 * time.Second
}

func (s *Stage) Initialize(dctx discovery.Context) error {
	if dctx.Platform() == nil {
		return fmt.Errorf("platform abstraction layer is not initialized in context")
	}
	if dctx.Platform().Network() == nil {
		return fmt.Errorf("network provider is not available for this platform")
	}
	return nil
}

func (s *Stage) Run(ctx context.Context, dctx discovery.Context) (discovery.DiscoveryArtifact, error) {
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

func (s *Stage) Validate(artifact discovery.DiscoveryArtifact) error {
	netInfo, ok := artifact.(models.NetworkInfo)
	if !ok {
		return fmt.Errorf("expected models.NetworkInfo artifact, got %T", artifact)
	}

	if len(netInfo.Interfaces) == 0 {
		return fmt.Errorf("invalid artifact: missing network interfaces")
	}

	return nil
}

func (s *Stage) Cleanup(ctx context.Context) error {
	// Nothing to clean up for network discovery
	return nil
}
