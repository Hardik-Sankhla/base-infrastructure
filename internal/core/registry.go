package core

import (
	"fmt"
	"sort"
	"sync"
)

// Registry manages stage registration and lookup. It prevents duplicate
// stage names and serves as a simple stage store.
type Registry struct {
	stages map[string]Stage
	mu     sync.RWMutex
}

// NewRegistry creates an empty stage registry.
func NewRegistry() *Registry {
	return &Registry{
		stages: make(map[string]Stage),
	}
}

// Register adds a stage to the registry. Returns an error if a stage
// with the same name is already registered.
func (r *Registry) Register(stage Stage) error {
	if stage == nil {
		return fmt.Errorf("cannot register nil stage")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	name := stage.Name()
	if _, exists := r.stages[name]; exists {
		return fmt.Errorf("stage %q is already registered", name)
	}

	r.stages[name] = stage
	return nil
}

// Get retrieves a stage by name. Returns the stage and true if found,
// or nil and false otherwise.
func (r *Registry) Get(name string) (Stage, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	s, ok := r.stages[name]
	return s, ok
}

// All returns all registered stages. The order is intentionally
// non-deterministic; the Pipeline is responsible for priority sorting.
func (r *Registry) All() []Stage {
	r.mu.RLock()
	defer r.mu.RUnlock()

	stages := make([]Stage, 0, len(r.stages))
	for _, s := range r.stages {
		stages = append(stages, s)
	}

	return stages
}

// Names returns the names of all registered stages, sorted alphabetically.
func (r *Registry) Names() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	names := make([]string, 0, len(r.stages))
	for name := range r.stages {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Len returns the number of registered stages.
func (r *Registry) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.stages)
}
