package storage

import (
	"testing"

	"fmt"

	"github.com/fitiavana07/tsk/pkg/storage/db"
	"github.com/fitiavana07/tsk/pkg/task"
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

func TestFileStorageTasks(t *testing.T) {
	file := tempFile(t)

	t.Run("NoTask", func(t *testing.T) {
		fs := NewFileStorage(file)
		l := len(fs.Tasks())
		if l != 0 {
			t.Errorf("got l=%d, want l=0", l)
		}
	})
	t.Run("OneTask", func(t *testing.T) {
		fs := FileStorage{file, db.DB{
			LastID: 1,
			Tasks: []task.Task{
				task.Task{
					ID:   1,
					Name: "a sample task",
				},
			},
		}}
		fs.Save()
		l := len(fs.Tasks())
		if l != 1 {
			t.Errorf("got l=%d, want l=1", l)
		}
	})
}

func TestFileStorageDoTask(t *testing.T) {
	file := tempFile(t)
	t.Run("DoTask", func(t *testing.T) {
		fs := FileStorage{file, db.DB{
			LastID: 1,
			Tasks: []task.Task{
				task.Task{
					ID:   1,
					Name: "a sample task",
				},
			},
		}}
		fs.Save()

		tsk := fs.DoTask(1)

		fmt.Printf("fs addr: %p\n", &fs)
		fmt.Printf("db addr: %p\n", &fs.db)
		for _, tsk := range fs.db.Tasks {
			fmt.Printf("task addr: %p\n", &tsk)
		}

		if tsk.ID != 1 {
			// t.Errorf("got wrong task: %s", tsk)
		}

		if tsk.State != task.StateDoing {
			t.Errorf("got returned state=%s, want state=%s", tsk.State, task.StateDoing)
		}

		got := fs.db.Tasks[0].State
		if got != task.StateDoing {
			t.Errorf("got original state=%s, want state=%s", got, task.StateDoing)
		}
	})

	t.Run("VerifyDoingTask", func(t *testing.T) {
		fs := NewFileStorage(file)
		tasks := fs.Tasks()
		tsk := tasks[0]

		if tsk.State != task.StateDoing {
			t.Errorf("state not persisted, got=%s, want=%s", tsk.State, task.StateDoing)
		}
	})
}
