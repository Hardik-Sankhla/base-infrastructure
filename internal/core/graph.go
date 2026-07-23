package core

import (
	"errors"
	"fmt"
	"sort"
)

// ExecutionGraph represents a Directed Acyclic Graph (DAG) of Tasks.
type ExecutionGraph struct {
	tasks    map[string]Task
	edges    map[string][]string // taskID -> list of taskIDs that depend on it
	inDegree map[string]int      // taskID -> number of dependencies it waits on
}

// NewExecutionGraph creates a new empty DAG.
func NewExecutionGraph() *ExecutionGraph {
	return &ExecutionGraph{
		tasks:    make(map[string]Task),
		edges:    make(map[string][]string),
		inDegree: make(map[string]int),
	}
}

// AddTask adds a task to the graph. Returns an error if the ID already exists.
func (g *ExecutionGraph) AddTask(t Task) error {
	id := t.ID()
	if _, exists := g.tasks[id]; exists {
		return fmt.Errorf("duplicate task ID: %s", id)
	}
	g.tasks[id] = t
	g.edges[id] = []string{}
	g.inDegree[id] = 0
	return nil
}

// AddDependency specifies that `fromID` must complete before `toID` can start.
// Returns an error if either node doesn't exist.
func (g *ExecutionGraph) AddDependency(fromID, toID string) error {
	if _, exists := g.tasks[fromID]; !exists {
		return fmt.Errorf("missing dependency node: %s", fromID)
	}
	if _, exists := g.tasks[toID]; !exists {
		return fmt.Errorf("missing dependency node: %s", toID)
	}

	g.edges[fromID] = append(g.edges[fromID], toID)
	g.inDegree[toID]++
	return nil
}

// TopologicalSort performs Kahn's algorithm to resolve task order.
// It returns tasks grouped into concurrent batches, or an error if a cycle is detected.
// To ensure deterministic ordering, nodes within the same batch are sorted alphabetically by ID.
func (g *ExecutionGraph) TopologicalSort() ([][]Task, error) {
	var batches [][]Task
	inDegree := make(map[string]int)
	for k, v := range g.inDegree {
		inDegree[k] = v
	}

	// Find initial zero in-degree nodes
	var queue []string
	for id, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, id)
		}
	}

	processedCount := 0

	for len(queue) > 0 {
		// Sort queue alphabetically to ensure deterministic topological ordering
		sort.Strings(queue)

		var currentBatch []Task
		var nextQueue []string

		for _, id := range queue {
			currentBatch = append(currentBatch, g.tasks[id])
			processedCount++

			// Decrease in-degree for dependents
			for _, depID := range g.edges[id] {
				inDegree[depID]--
				if inDegree[depID] == 0 {
					nextQueue = append(nextQueue, depID)
				}
			}
		}

		batches = append(batches, currentBatch)
		queue = nextQueue
	}

	if processedCount != len(g.tasks) {
		return nil, errors.New("circular dependency detected in execution graph")
	}

	return batches, nil
}
