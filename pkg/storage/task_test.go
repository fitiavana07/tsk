package storage

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	file := tempFile(t)
	fs := NewFileStorage(file)

	t.Run("AddFirstTask", func(t *testing.T) {
		taskName := "do the task"
		task := fs.AddTask(taskName)
		if task.ID != 1 {
			t.Errorf("wrong task id, got %d, want 1", task.ID)
		}
		if task.Name != taskName {
			t.Errorf("wrong task name, got %q, want %q", task.Name, taskName)
		}
	})
}
