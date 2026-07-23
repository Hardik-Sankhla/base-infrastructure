package state

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/base-infrastructure/platform/internal/domain/models"
)

func TestParser_Parse(t *testing.T) {
	yamlData := []byte(`
version: "1.0"
settings:
  rollback_on_failure: true
  strict_version_match: false
capabilities:
  - id: "docker"
    provider: "system"
    version: ">=20.10.0"
    state: "available"
    properties:
      socket: "/var/run/docker.sock"
`)

	parser := NewParser()
	manifest, err := parser.Parse(yamlData)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if manifest.Version != "1.0" {
		t.Errorf("expected version 1.0, got %s", manifest.Version)
	}

	if !manifest.Settings.RollbackOnFailure {
		t.Errorf("expected RollbackOnFailure true, got false")
	}

	if len(manifest.Capabilities) != 1 {
		t.Fatalf("expected 1 capability, got %d", len(manifest.Capabilities))
	}

	cap := manifest.Capabilities[0]
	if cap.ID != "docker" {
		t.Errorf("expected capability ID 'docker', got %s", cap.ID)
	}
	if cap.State != models.StateAvailable {
		t.Errorf("expected capability state 'available', got %s", cap.State)
	}
	if cap.Properties["socket"] != "/var/run/docker.sock" {
		t.Errorf("expected socket property '/var/run/docker.sock', got %s", cap.Properties["socket"])
	}
}

func TestParser_Load(t *testing.T) {
	yamlData := []byte(`
version: "1.0"
capabilities:
  - id: "git"
    state: "available"
`)
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "state.yaml")
	if err := os.WriteFile(filePath, yamlData, 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	parser := NewParser()
	manifest, err := parser.Load(filePath)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if manifest.Version != "1.0" {
		t.Errorf("expected version 1.0, got %s", manifest.Version)
	}
	if len(manifest.Capabilities) != 1 {
		t.Fatalf("expected 1 capability, got %d", len(manifest.Capabilities))
	}
}

func TestParser_ParseInvalid(t *testing.T) {
	yamlData := []byte(`
version: "1.0"
capabilities:
  - id: "git"
    state: [invalid, format]
`)
	parser := NewParser()
	_, err := parser.Parse(yamlData)
	if err == nil {
		t.Errorf("expected error for invalid YAML, got nil")
	}
}
