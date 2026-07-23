package network

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/base-infrastructure/platform/internal/domain/models"
	"github.com/base-infrastructure/platform/internal/platform/mock"
	"github.com/base-infrastructure/platform/internal/runtime"

	dctx "github.com/base-infrastructure/platform/internal/discovery"
)

func TestNetworkStage_Success(t *testing.T) {
	stage := NewStage()
	p := mock.NewPlatform()

	p.MockNetwork.Interfaces = []models.NetworkInterface{
		{
			Name: "eth0",
			MAC:  "00:11:22:33:44:55",
			MTU:  1500,
			IPv4: []string{"192.168.1.10"},
			IsUp: true,
		},
	}
	p.MockNetwork.DNS = models.DNSConfig{
		Servers: []string{"8.8.8.8"},
	}

	log := slog.Default()
	bus := runtime.NewEventBus()
	ctx := dctx.NewContext(log, bus, nil, nil, p)

	if err := stage.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize: %v", err)
	}

	goCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	artifact, err := stage.Run(goCtx, ctx)
	if err != nil {
		t.Fatalf("Stage run failed: %v", err)
	}

	if err := stage.Validate(artifact); err != nil {
		t.Fatalf("Artifact validation failed: %v", err)
	}

	netInfo, ok := artifact.(models.NetworkInfo)
	if !ok {
		t.Fatalf("Expected models.NetworkInfo, got %T", artifact)
	}

	if len(netInfo.Interfaces) != 1 || netInfo.Interfaces[0].Name != "eth0" {
		t.Errorf("Unexpected interfaces: %+v", netInfo.Interfaces)
	}
	if len(netInfo.DNS.Servers) != 1 || netInfo.DNS.Servers[0] != "8.8.8.8" {
		t.Errorf("Unexpected DNS config: %+v", netInfo.DNS)
	}
}

func TestNetworkStage_UninitializedPlatform(t *testing.T) {
	stage := NewStage()
	log := slog.Default()
	bus := runtime.NewEventBus()

	// Context with NO platform
	ctx := dctx.NewContext(log, bus, nil, nil, nil)

	err := stage.Initialize(ctx)
	if err == nil {
		t.Fatal("Expected initialization to fail with no platform")
	}
}

func TestNetworkStage_GetInterfacesError(t *testing.T) {
	stage := NewStage()
	p := mock.NewPlatform()

	p.MockNetwork.Err = errors.New("network provider error")

	log := slog.Default()
	bus := runtime.NewEventBus()
	ctx := dctx.NewContext(log, bus, nil, nil, p)

	if err := stage.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize: %v", err)
	}

	goCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := stage.Run(goCtx, ctx)
	if err == nil {
		t.Fatal("Expected run to fail when interfaces error occurs")
	}
}
