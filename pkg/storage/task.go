package storage

// storage/task.go contains task management features provided by the package storage.

import (
	"github.com/fitiavana07/tsk/pkg/task"
)

// AddTask to the storage in-memory representation
func (fs FileStorage) AddTask(name string) *task.Task {
	tsk := fs.db.AddTask(name)
	fs.Save()
	return tsk
}

// Tasks returns a list of Todo and Doing tasks stored in the FileStorage.
// Done tasks are not part of the returned list.
// To get list of done tasks, use FileStorage.TasksDone().
func (fs FileStorage) Tasks() (r []task.Task) {
	for _, t := range fs.db.Tasks {
		if t.State == task.StateTodo || t.State == task.StateDoing {
			r = append(r, t)
		}
	}
	return
}

// TasksDone returns a list of all done tasks.
func (fs FileStorage) TasksDone() (r []task.Task) {
	for _, t := range fs.db.Tasks {
		if t.State == task.StateDone {
			r = append(r, t)
		}
	}
	return
}

// DoTask moves the state of a task from Todo to Doing.
func (fs *FileStorage) DoTask(id int) *task.Task {
	for i, tsk := range fs.db.Tasks {
		if tsk.ID == id {
			fs.db.Tasks[i].State = task.StateDoing
			fs.Save()
			return &fs.db.Tasks[i]
		}
	}
	return nil
}

// DoneTask moves the state of a task from Doing to Done.
func (fs *FileStorage) DoneTask(id int) *task.Task {
	for i, tsk := range fs.db.Tasks {
		if tsk.ID == id {
			fs.db.Tasks[i].State = task.StateDone
			fs.Save()
			return &fs.db.Tasks[i]
		}
	}
	return nil
}
