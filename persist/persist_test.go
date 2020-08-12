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

	// TODO mock ioutil.WriteFile
	var persister Persister
	persister = TskDataPersister{}

	// act
	persister.Write(tskData)

	// check
	// TODO check if the file was written properly
	// TODO check if the written data match the original data
}
