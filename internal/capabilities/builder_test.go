package capabilities

import (
	"testing"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

func TestBuilder_Build(t *testing.T) {
	manifest := &models.DiscoveryManifest{
		Artifacts: map[string]any{
			"network": models.NetworkInfo{
				Interfaces: []models.NetworkInterface{
					{
						Name: "eth0",
						IsUp: true,
						IPv4: []string{"192.168.1.10"},
					},
				},
			},
			"software": models.SoftwareInfo{
				Runtimes: []models.RuntimeEnvironment{
					{
						Name:    "docker",
						Version: "24.0.5",
						Path:    "/usr/bin/docker",
					},
				},
			},
		},
	}

	builder := NewBuilder(manifest)
	caps := builder.Build()

	if len(caps) != 3 {
		t.Fatalf("Expected 3 capabilities, got %d", len(caps))
	}

	foundNetwork := false
	foundDocker := false
	foundContainerRuntime := false

	for _, c := range caps {
		switch c.ID {
		case "network.connectivity":
			foundNetwork = true
		case "runtime.docker":
			foundDocker = true
			if c.Version != "24.0.5" {
				t.Errorf("Expected docker version 24.0.5, got %s", c.Version)
			}
		case "container.runtime":
			foundContainerRuntime = true
			if c.Provider != "docker" {
				t.Errorf("Expected container runtime provider docker, got %s", c.Provider)
			}
		}
	}

	if !foundNetwork {
		t.Error("Did not find network.connectivity capability")
	}
	if !foundDocker {
		t.Error("Did not find runtime.docker capability")
	}
	if !foundContainerRuntime {
		t.Error("Did not find container.runtime capability")
	}
}

func TestBuilder_Build_Empty(t *testing.T) {
	manifest := &models.DiscoveryManifest{}
	builder := NewBuilder(manifest)
	caps := builder.Build()

	if len(caps) != 0 {
		t.Fatalf("Expected 0 capabilities, got %d", len(caps))
	}
}
