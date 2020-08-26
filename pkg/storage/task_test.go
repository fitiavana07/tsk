package storage

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	file := tempFile(t)

	t.Run("AddFirstTask", func(t *testing.T) {
		fs := NewFileStorage(file)
		taskName := "do the task"
		task := fs.AddTask(taskName)
		if task.ID != 1 {
			t.Errorf("wrong task id, got %d, want 1", task.ID)
		}
		if task.Name != taskName {
			t.Errorf("wrong task name, got %q, want %q", task.Name, taskName)
		}
	})

	t.Run("AddSecondTask", func(t *testing.T) {
		fs := NewFileStorage(file)
		taskName := "do the 2nd task"
		task := fs.AddTask(taskName)
		if task.ID != 2 {
			t.Errorf("wrong task id, got=%d, want=2", task.ID)
		}
		if task.Name != taskName {
			t.Errorf("wrong task name, got %q, want %q", task.Name, taskName)
		}
	})
}
