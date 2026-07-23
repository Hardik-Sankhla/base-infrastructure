package core

import "context"

// Task represents a discrete, context-aware unit of work in the ExecutionGraph.
type Task interface {
	// ID returns a unique identifier for this task within the graph.
	ID() string

	// CheckIdempotency verifies if the task's desired state is already met.
	// Returns true if the system is already in the desired state, false otherwise.
	CheckIdempotency(ctx context.Context) (bool, error)

	// Execute applies the state change.
	Execute(ctx context.Context) error

	// Rollback reverses the state change safely if the pipeline fails.
	Rollback(ctx context.Context) error
}
