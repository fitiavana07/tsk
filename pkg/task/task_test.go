package task

import (
	"testing"
)

/*
Testing of domain of the application.

A task is represented by a value of type Task.
A Task has these attributes:
	- an ID: an integer: 1, 2, 3 ...
	- a name: a string
	- a state: one of Todo, Doing, Done (Future: Archived, Trashed)

	(future, still to decide about, inspire from GTD)
	- a date
	- a context
*/

func TestNew(t *testing.T) {
	taskName := "do the task"
	task := New(taskName)

	t.Run("CheckId", func(t *testing.T) {
		got := task.ID
		if got != 0 {
			t.Errorf("got task.ID=%d, want 0", got)
		}
	})
	t.Run("CheckName", func(t *testing.T) {
		got := task.Name
		if got != taskName {
			t.Errorf("got task.Name=%q, want=%q", got, taskName)
		}
	})
}
