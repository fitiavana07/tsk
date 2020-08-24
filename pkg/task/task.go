package task

// Task is something to do, moved into doing, then marked done
type Task struct {
	id    int
	name  string
	state State
}

// New constructs a new task
func New() *Task {
	return &Task{}
}

// ID returns the id of a task
func (t *Task) ID() int {
	return t.id
}

// State is the state a task is in
type State uint8

// Possible states of tasks
const (
	StateTodo State = iota
	StateDoing
	StateDone
)
