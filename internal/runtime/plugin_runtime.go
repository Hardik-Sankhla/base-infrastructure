package runtime

import (
	"context"

	"fmt"
	"log/slog"
)

// Result represents the standardized response from a plugin
type Result struct {
	Success bool
	Data    interface{}
	Error   string
}

// Interface defines the strict contract all plugins must implement
type PluginInterface interface {
	Discover(ctx context.Context) (Result, error)
	Plan(ctx context.Context) (Result, error)
	Install(ctx context.Context) (Result, error)
	Configure(ctx context.Context) (Result, error)
	Verify(ctx context.Context) (Result, error)
	Health(ctx context.Context) (Result, error)
	Update(ctx context.Context) (Result, error)
	Rollback(ctx context.Context) (Result, error)
	Uninstall(ctx context.Context) (Result, error)
	Cleanup(ctx context.Context) (Result, error)
}

type PluginRegistry interface {
	Register(manifest *PluginManifest, instance Interface) error
	Get(name string) (Interface, error)
}

type DefaultPluginRegistry struct {
	plugins map[string]Interface
}

func NewPluginRegistry() *DefaultPluginRegistry {
	return &DefaultRegistry{
		plugins: make(map[string]Interface),
	}
}

func (r *DefaultPluginRegistry) Register(manifest *PluginManifest, instance Interface) error {
	if _, exists := r.plugins[manifest.Name]; exists {
		return fmt.Errorf("plugin %s is already registered", manifest.Name)
	}
	r.plugins[manifest.Name] = instance
	slog.Info("Registered plugin", "name", manifest.Name, "version", manifest.Version)
	return nil
}

func (r *DefaultPluginRegistry) Get(name string) (Interface, error) {
	if instance, exists := r.plugins[name]; exists {
		return instance, nil
	}
	return nil, fmt.Errorf("plugin %s not found", name)
}
