package core

import (
	"context"
	"fmt"
)

// EngineOptions configures execution behavior.
type EngineOptions struct {
	RollbackOnFailure bool
	DryRun            bool
}

// StateEngine orchestrates the safe execution of an ExecutionGraph.
type StateEngine struct {
	options EngineOptions
}

// NewStateEngine creates a new state execution engine.
func NewStateEngine(opts EngineOptions) *StateEngine {
	return &StateEngine{
		options: opts,
	}
}

// Execute resolves and applies the graph safely.
func (e *StateEngine) Execute(ctx context.Context, graph *ExecutionGraph) error {
	batches, err := graph.TopologicalSort()
	if err != nil {
		return fmt.Errorf("graph validation failed: %w", err)
	}

	var executedTasks []Task

	// Execute sequentially for Sprint 2
	for _, batch := range batches {
		for _, task := range batch {
			id := task.ID()

			// Check Idempotency
			isIdempotent, err := task.CheckIdempotency(ctx)
			if err != nil {
				e.handleRollback(ctx, executedTasks)
				return fmt.Errorf("task [%s] idempotency check failed: %w", id, err)
			}

			if isIdempotent {
				// Task is already satisfied, skip execution
				continue
			}

			// Dry Run check
			if e.options.DryRun {
				continue
			}

			// Execute
			err = task.Execute(ctx)
			if err != nil {
				e.handleRollback(ctx, executedTasks)
				return fmt.Errorf("task [%s] execution failed: %w", id, err)
			}

			// Record successfully executed task for potential rollback
			executedTasks = append(executedTasks, task)
		}
	}

	return nil
}

// handleRollback walks executed tasks in reverse order to undo changes.
func (e *StateEngine) handleRollback(ctx context.Context, executedTasks []Task) {
	if !e.options.RollbackOnFailure {
		return
	}

	// Reverse iteration
	for i := len(executedTasks) - 1; i >= 0; i-- {
		task := executedTasks[i]
		// In a production system, we might collect rollback errors without halting
		// the rollback process for other tasks.
		// For now, we attempt rollback and ignore errors since the primary error
		// from execution has already been returned.
		_ = task.Rollback(ctx)
	}
}
