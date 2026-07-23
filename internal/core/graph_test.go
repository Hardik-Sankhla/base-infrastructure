package core

import (
	"context"
	"reflect"
	"testing"
)

// mockTask implements Task for testing
type mockTask struct {
	id string
}

func (m *mockTask) ID() string                                         { return m.id }
func (m *mockTask) CheckIdempotency(ctx context.Context) (bool, error) { return false, nil }
func (m *mockTask) Execute(ctx context.Context) error                  { return nil }
func (m *mockTask) Rollback(ctx context.Context) error                 { return nil }

func TestExecutionGraph(t *testing.T) {
	t.Run("Empty graph", func(t *testing.T) {
		g := NewExecutionGraph()
		batches, err := g.TopologicalSort()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(batches) != 0 {
			t.Errorf("expected 0 batches, got %d", len(batches))
		}
	})

	t.Run("Single task", func(t *testing.T) {
		g := NewExecutionGraph()
		_ = g.AddTask(&mockTask{id: "A"})
		batches, err := g.TopologicalSort()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(batches) != 1 {
			t.Fatalf("expected 1 batch, got %d", len(batches))
		}
		if batches[0][0].ID() != "A" {
			t.Errorf("expected task A, got %s", batches[0][0].ID())
		}
	})

	t.Run("Duplicate task IDs", func(t *testing.T) {
		g := NewExecutionGraph()
		err1 := g.AddTask(&mockTask{id: "A"})
		err2 := g.AddTask(&mockTask{id: "A"})
		if err1 != nil {
			t.Fatalf("unexpected error on first add: %v", err1)
		}
		if err2 == nil {
			t.Error("expected error on duplicate add, got nil")
		}
	})

	t.Run("Missing dependency", func(t *testing.T) {
		g := NewExecutionGraph()
		_ = g.AddTask(&mockTask{id: "A"})
		err := g.AddDependency("A", "B")
		if err == nil {
			t.Error("expected error on missing dependency, got nil")
		}
	})

	t.Run("Dependency ordering", func(t *testing.T) {
		g := NewExecutionGraph()
		_ = g.AddTask(&mockTask{id: "A"})
		_ = g.AddTask(&mockTask{id: "B"})
		_ = g.AddTask(&mockTask{id: "C"})

		_ = g.AddDependency("A", "B") // A -> B
		_ = g.AddDependency("B", "C") // B -> C

		batches, err := g.TopologicalSort()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(batches) != 3 {
			t.Fatalf("expected 3 batches, got %d", len(batches))
		}
		if batches[0][0].ID() != "A" {
			t.Errorf("batch 0: expected A, got %s", batches[0][0].ID())
		}
		if batches[1][0].ID() != "B" {
			t.Errorf("batch 1: expected B, got %s", batches[1][0].ID())
		}
		if batches[2][0].ID() != "C" {
			t.Errorf("batch 2: expected C, got %s", batches[2][0].ID())
		}
	})

	t.Run("Deterministic ordering", func(t *testing.T) {
		g := NewExecutionGraph()
		_ = g.AddTask(&mockTask{id: "C"})
		_ = g.AddTask(&mockTask{id: "A"})
		_ = g.AddTask(&mockTask{id: "B"})

		batches, err := g.TopologicalSort()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(batches) != 1 {
			t.Fatalf("expected 1 batch, got %d", len(batches))
		}

		var ids []string
		for _, task := range batches[0] {
			ids = append(ids, task.ID())
		}

		expected := []string{"A", "B", "C"}
		if !reflect.DeepEqual(ids, expected) {
			t.Errorf("expected deterministic order %v, got %v", expected, ids)
		}
	})

	t.Run("Circular dependency", func(t *testing.T) {
		g := NewExecutionGraph()
		_ = g.AddTask(&mockTask{id: "A"})
		_ = g.AddTask(&mockTask{id: "B"})

		_ = g.AddDependency("A", "B")
		_ = g.AddDependency("B", "A")

		_, err := g.TopologicalSort()
		if err == nil {
			t.Error("expected error on circular dependency, got nil")
		}
	})
}
