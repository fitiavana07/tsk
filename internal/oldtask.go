package internal

// State represents the type of a task
// (define a new type to allow associating methods)
type State int

// States of Tasks
const (
	// todo state
	StateTodo State = iota

	// doing state
	StateDoing

	// done state
	StateDone
)

// Task type represents a task, with an index, a name and a state
type Task struct {

	// index: 1, 2, 3
	Index int

	// name of the task
	Name string

	// state of the task
	State State
}
