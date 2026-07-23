package discovery

import (
	"testing"
	"time"

	"github.com/base-infrastructure/platform/internal/runtime"
)

type mockStage struct {
	name      string
	dependsOn []string
}

func (m *mockStage) Name() string                  { return m.name }
func (m *mockStage) Version() string               { return "1.0" }
func (m *mockStage) Description() string           { return "mock" }
func (m *mockStage) Priority() int                 { return 0 }
func (m *mockStage) DependsOn() []string           { return m.dependsOn }
func (m *mockStage) Timeout() time.Duration        { return time.Second }
func (m *mockStage) Initialize(dctx Context) error { return nil }
func (m *mockStage) Run(ctx runtime.Context, dctx Context) (DiscoveryArtifact, error) {
	return nil, nil
}
func (m *mockStage) Validate(artifact DiscoveryArtifact) error { return nil }
func (m *mockStage) Cleanup(ctx runtime.Context) error         { return nil }

func TestValidator_ValidGraph(t *testing.T) {
	v := NewValidator()
	stages := []Stage{
		&mockStage{name: "a", dependsOn: nil},
		&mockStage{name: "b", dependsOn: []string{"a"}},
		&mockStage{name: "c", dependsOn: []string{"a", "b"}},
	}
	if err := v.Validate(stages); err != nil {
		t.Fatalf("expected valid graph, got error: %v", err)
	}
}

func TestValidator_DuplicateStage(t *testing.T) {
	v := NewValidator()
	stages := []Stage{
		&mockStage{name: "a", dependsOn: nil},
		&mockStage{name: "a", dependsOn: nil},
	}
	if err := v.Validate(stages); err == nil {
		t.Fatalf("expected duplicate stage error")
	}
}

func TestValidator_MissingDependency(t *testing.T) {
	v := NewValidator()
	stages := []Stage{
		&mockStage{name: "a", dependsOn: []string{"missing"}},
	}
	if err := v.Validate(stages); err == nil {
		t.Fatalf("expected missing dependency error")
	}
}

func TestValidator_CircularDependency(t *testing.T) {
	v := NewValidator()
	stages := []Stage{
		&mockStage{name: "a", dependsOn: []string{"b"}},
		&mockStage{name: "b", dependsOn: []string{"c"}},
		&mockStage{name: "c", dependsOn: []string{"a"}},
	}
	if err := v.Validate(stages); err == nil {
		t.Fatalf("expected circular dependency error")
	}
}
