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

// Tasks returns a list of all tasks stored in the FileStorage.
func (fs FileStorage) Tasks() []task.Task {
	return fs.db.Tasks
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
