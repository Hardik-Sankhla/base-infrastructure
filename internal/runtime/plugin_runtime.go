package runtime

import (
	"context"
	"fmt"
	"log/slog"
)

// PluginResult represents the standardized response from a plugin
type PluginResult struct {
	Success bool
	Data    interface{}
	Error   string
}

// Interface defines the strict contract all plugins must implement
type PluginInterface interface {
	Discover(ctx context.Context) (PluginResult, error)
	Plan(ctx context.Context) (PluginResult, error)
	Install(ctx context.Context) (PluginResult, error)
	Configure(ctx context.Context) (PluginResult, error)
	Verify(ctx context.Context) (PluginResult, error)
	Health(ctx context.Context) (PluginResult, error)
	Update(ctx context.Context) (PluginResult, error)
	Rollback(ctx context.Context) (PluginResult, error)
	Uninstall(ctx context.Context) (PluginResult, error)
	Cleanup(ctx context.Context) (PluginResult, error)
}

type PluginRegistry interface {
	Register(manifest *PluginManifest, instance PluginInterface) error
	Get(name string) (PluginInterface, error)
}

type DefaultPluginRegistry struct {
	plugins map[string]PluginInterface
}

func NewPluginRegistry() *DefaultPluginRegistry {
	return &DefaultPluginRegistry{
		plugins: make(map[string]PluginInterface),
	}
}

func (r *DefaultPluginRegistry) Register(manifest *PluginManifest, instance PluginInterface) error {
	if _, exists := r.plugins[manifest.Name]; exists {
		return fmt.Errorf("plugin %s is already registered", manifest.Name)
	}
	r.plugins[manifest.Name] = instance
	slog.Info("Registered plugin", "name", manifest.Name, "version", manifest.Version)
	return nil
}

func (r *DefaultPluginRegistry) Get(name string) (PluginInterface, error) {
	if instance, exists := r.plugins[name]; exists {
		return instance, nil
	}
	return nil, fmt.Errorf("plugin %s not found", name)
}
