package persist

import (
	// "fmt"
	"testing"

	"github.com/fitiavana07/tsk/task"
)

func TestFilePersisterWrite(t *testing.T) {
	// data
	tskData := TskData{
		LastTaskIndex: 1,
		Tasks: []task.Task{
			{
				Index: 1,
				Name:  "do this",
				State: task.StateDone,
			},
		},
	}
	var persister Persister
	persister = TskDataPersister{tskData}
	persister.Write()
}
