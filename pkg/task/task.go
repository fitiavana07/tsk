package task

import (
	"fmt"
)

// Task is something to do, moved into doing, then marked done
type Task struct {
	ID    int
	Name  string
	State State
}

// New constructs a new task
func New(id int, name string) *Task {
	return &Task{ID: id, Name: name}
}

func (t Task) String() string {
	return fmt.Sprintf(
		"Task(ID=%d, Name=%q, State=%s)",
		t.ID,
		t.Name,
		t.State,
	)
}

// State is the state a task is in
type State uint8

func (st State) String() string {
	switch st {
	case StateTodo:
		return "Todo"
	case StateDoing:
		return "Doing"
	case StateDone:
		return "Done"
	default:
		return ""
	}
}

// Possible states of tasks
const (
	StateTodo State = iota
	StateDoing
	StateDone
)
