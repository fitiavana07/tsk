package task

// Task type represents a task, with an index and a name
type Task struct {
	Index int
	Name  string
	State State
}

// State is the state of a task
type State = uint8

// Possible states of tasks
const (
	StateTodo State = iota
	StateDoing
	StateDone
)
