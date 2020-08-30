package db

import (
	"github.com/fitiavana07/tsk/pkg/task"
)

// DB is the database in file.
type DB struct {
	LastID int
	Tasks  []task.Task
}

// AddTask add a task using its name  to the db.
func (db *DB) AddTask(name string) *task.Task {
	newID := db.LastID + 1

	tsk := task.New(newID, name)
	db.Tasks = append(db.Tasks, *tsk)

	db.LastID = newID
	return tsk
}
