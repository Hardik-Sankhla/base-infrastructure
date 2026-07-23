package runtime

import (
	"context"
	"fmt"
	"log/slog"
)

// Status represents the current state of a task
type Status string

const (
	Pending    Status = "Pending"
	Running    Status = "Running"
	Completed  Status = "Completed"
	Failed     Status = "Failed"
	RolledBack Status = "RolledBack"
)

// Task interface
type Task interface {
	Name() string
	Execute(ctx context.Context) error
	Rollback(ctx context.Context) error
}

// TaskEngine manages execution of tasks
type TaskEngine interface {
	Submit(ctx context.Context, task Task) error
}

type DefaultTaskEngine struct{}

func NewTaskEngine() *DefaultTaskEngine {
	return &DefaultTaskEngine{}
}

func (e *DefaultTaskEngine) Submit(ctx context.Context, task Task) error {
	slog.Info("Executing task", "task", task.Name())

	err := task.Execute(ctx)
	if err != nil {
		slog.Error("Task failed", "task", task.Name(), "error", err)

		slog.Info("Attempting rollback", "task", task.Name())
		if rollbackErr := task.Rollback(ctx); rollbackErr != nil {
			slog.Error("Rollback failed", "task", task.Name(), "error", rollbackErr)
			return fmt.Errorf("task execution failed: %v, rollback also failed: %v", err, rollbackErr)
		}
		return fmt.Errorf("task failed but successfully rolled back: %w", err)
	}

	slog.Info("Task completed successfully", "task", task.Name())
	return nil
}
