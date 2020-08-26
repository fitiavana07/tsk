package task

// Task is something to do, moved into doing, then marked done
type Task struct {
	ID    int
	Name  string
	state State
}

// New constructs a new task
func New(name string) *Task {
	return &Task{Name: name}
}

// State is the state a task is in
type State uint8

// Possible states of tasks
const (
	StateTodo State = iota
	StateDoing
	StateDone
)
