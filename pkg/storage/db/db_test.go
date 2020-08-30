package db

import (
	"testing"
)

func TestDBAddTask(t *testing.T) {
	db := DB{}

	t.Run("FirstAdd", func(t *testing.T) {
		taskName := "do this task"
		tsk := db.AddTask(taskName)

		if tsk.ID != 1 {
			t.Errorf("got tsk.ID=%d, want 1", tsk.ID)
		}

		found := false
		for _, tk := range db.Tasks {
			if tk.Name == taskName && tk.ID == tsk.ID {
				found = true
			}
		}
		if !found {
			t.Errorf("added task not found in db")
		}
	})
	t.Run("SecondAdd", func(t *testing.T) {
		taskName := "do this  second task"
		tsk := db.AddTask(taskName)

		if tsk.ID != 2 {
			t.Errorf("got tsk.ID=%d, want 1", tsk.ID)
		}

		found := false
		for _, tk := range db.Tasks {
			if tk.Name == taskName && tk.ID == tsk.ID {
				found = true
			}
		}
		if !found {
			t.Errorf("added task not found in db")
		}
	})
}
