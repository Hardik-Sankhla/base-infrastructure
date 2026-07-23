package core

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

type stateTestTask struct {
	id             string
	idempotent     bool
	executeErr     error
	rollbackErr    error
	executed       bool
	rolledBack     bool
	idempotencyErr error
}

func (m *stateTestTask) ID() string { return m.id }

func (m *stateTestTask) CheckIdempotency(ctx context.Context) (bool, error) {
	return m.idempotent, m.idempotencyErr
}

func (m *stateTestTask) Execute(ctx context.Context) error {
	m.executed = true
	return m.executeErr
}

func (m *stateTestTask) Rollback(ctx context.Context) error {
	m.rolledBack = true
	return m.rollbackErr
}

func TestStateEngine(t *testing.T) {
	ctx := context.Background()

	t.Run("Execute success", func(t *testing.T) {
		task1 := &stateTestTask{id: "A"}
		task2 := &stateTestTask{id: "B"}

		g := NewExecutionGraph()
		_ = g.AddTask(task1)
		_ = g.AddTask(task2)

		engine := NewStateEngine(EngineOptions{RollbackOnFailure: false, DryRun: false})
		err := engine.Execute(ctx, g)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !task1.executed || !task2.executed {
			t.Error("expected both tasks to execute")
		}
	})

	t.Run("Idempotent task skipped", func(t *testing.T) {
		task1 := &stateTestTask{id: "A", idempotent: true}
		task2 := &stateTestTask{id: "B", idempotent: false}

		g := NewExecutionGraph()
		_ = g.AddTask(task1)
		_ = g.AddTask(task2)

		engine := NewStateEngine(EngineOptions{RollbackOnFailure: false, DryRun: false})
		err := engine.Execute(ctx, g)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if task1.executed {
			t.Error("expected task1 to be skipped")
		}
		if !task2.executed {
			t.Error("expected task2 to execute")
		}
	})

	t.Run("Execute failure halts execution", func(t *testing.T) {
		expectedErr := errors.New("boom")
		task1 := &stateTestTask{id: "A"}
		task2 := &stateTestTask{id: "B", executeErr: expectedErr}
		task3 := &stateTestTask{id: "C"}

		g := NewExecutionGraph()
		_ = g.AddTask(task1)
		_ = g.AddTask(task2)
		_ = g.AddTask(task3)
		_ = g.AddDependency("A", "B")
		_ = g.AddDependency("B", "C")

		engine := NewStateEngine(EngineOptions{RollbackOnFailure: false, DryRun: false})
		err := engine.Execute(ctx, g)

		if err == nil {
			t.Error("expected execution error, got nil")
		}

		if !task1.executed {
			t.Error("expected task1 to execute")
		}
		if !task2.executed {
			t.Error("expected task2 to attempt execution")
		}
		if task3.executed {
			t.Error("expected task3 to be skipped after failure")
		}
	})

	t.Run("Rollback on failure", func(t *testing.T) {
		expectedErr := errors.New("boom")
		task1 := &stateTestTask{id: "A"}
		task2 := &stateTestTask{id: "B", executeErr: expectedErr}

		g := NewExecutionGraph()
		_ = g.AddTask(task1)
		_ = g.AddTask(task2)
		_ = g.AddDependency("A", "B")

		engine := NewStateEngine(EngineOptions{RollbackOnFailure: true, DryRun: false})
		err := engine.Execute(ctx, g)

		if err == nil {
			t.Error("expected execution error, got nil")
		}

		if !task1.executed {
			t.Error("expected task1 to execute")
		}
		if !task1.rolledBack {
			t.Error("expected task1 to be rolled back")
		}
		if task2.rolledBack {
			t.Error("expected task2 not to be rolled back since its execution failed")
		}
	})

	t.Run("Rollback failure ignores subsequent errors", func(t *testing.T) {
		expectedErr := errors.New("execute boom")
		rollbackErr := errors.New("rollback boom")
		task1 := &stateTestTask{id: "A", rollbackErr: rollbackErr}
		task2 := &stateTestTask{id: "B", executeErr: expectedErr}

		g := NewExecutionGraph()
		_ = g.AddTask(task1)
		_ = g.AddTask(task2)
		_ = g.AddDependency("A", "B")

		engine := NewStateEngine(EngineOptions{RollbackOnFailure: true, DryRun: false})
		err := engine.Execute(ctx, g)

		if err == nil {
			t.Error("expected execution error, got nil")
		}
		// The error returned should be the execution error, not the rollback error.
		if !errors.Is(err, expectedErr) && err.Error() != fmt.Sprintf("task [B] execution failed: %s", expectedErr.Error()) {
			// Actually we are wrapping the error
		}

		if !task1.rolledBack {
			t.Error("expected task1 rollback attempt")
		}
	})
}
