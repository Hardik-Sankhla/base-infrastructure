package discovery

import (
	"fmt"
	"strings"
)

// Validator validates the graph of discovery stages before execution.
type Validator struct{}

// NewValidator creates a new dependency validator.
func NewValidator() *Validator {
	return &Validator{}
}

// Validate takes a list of stages and ensures they form a valid directed acyclic graph (DAG).
// It checks for:
// 1. Duplicate stage names.
// 2. Missing dependencies.
// 3. Circular dependencies.
func (v *Validator) Validate(stages []Stage) error {
	stageMap := make(map[string]Stage)
	for _, s := range stages {
		name := s.Name()
		if _, exists := stageMap[name]; exists {
			return fmt.Errorf("duplicate stage detected: %s", name)
		}
		stageMap[name] = s
	}

	// Check for missing dependencies
	for _, s := range stages {
		for _, dep := range s.DependsOn() {
			if _, exists := stageMap[dep]; !exists {
				return fmt.Errorf("stage '%s' depends on unknown stage '%s'", s.Name(), dep)
			}
		}
	}

	// Detect cycles using Depth-First Search (DFS)
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	var hasCycle func(name string, path []string) error
	hasCycle = func(name string, path []string) error {
		if recStack[name] {
			cyclePath := append(path, name)
			return fmt.Errorf("circular dependency detected: %s", strings.Join(cyclePath, " -> "))
		}
		if visited[name] {
			return nil
		}

		visited[name] = true
		recStack[name] = true
		path = append(path, name)

		stage := stageMap[name]
		for _, dep := range stage.DependsOn() {
			if err := hasCycle(dep, path); err != nil {
				return err
			}
		}

		recStack[name] = false
		return nil
	}

	for _, s := range stages {
		if !visited[s.Name()] {
			if err := hasCycle(s.Name(), nil); err != nil {
				return err
			}
		}
	}

	return nil
}
