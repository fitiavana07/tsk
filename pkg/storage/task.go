package storage

import (
	"github.com/fitiavana07/tsk/pkg/task"
)

// AddTask to the storage in-memory representation
func (fs FileStorage) AddTask(name string) *task.Task {
	// newID := fs.db.LastID + 1
	// tsk := task.New(newID, name)

	// append(fs.db.Tasks, tsk)
	// fs.db.LastID = newID

	tsk := task.New(1, name)
	return tsk
}
