package persist

import (
	"github.com/fitiavana07/tsk/task"
)

// file name
const DataFileName = ".tsk"

// TskData represents the whole tsk data
type TskData struct {
	// last task index, used to create the next task
	lastTaskIndex int

	// list of tasks
	tasks []task.Task
}

// TskDataStorage represents the storage of the data
