package storage

import (
	"github.com/fitiavana07/tsk/pkg/task"
)

// AddTask to the storage in-memory representation
func (fs FileStorage) AddTask(name string) *task.Task {
	tsk := fs.db.AddTask(name)
	fs.Save()
	return tsk
}
