package persist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fitiavana07/tsk/task"
)

// file name
const DefaultDataFileName = ".tsk"

// TskData represents the whole tsk data
type TskData struct {
	// last task index, used to create the next task
	LastTaskIndex int

	// list of tasks
	Tasks []task.Task
}

// Persister interface for persistance functions
type Persister interface {
	Read()
	Write()
}

// TskDataPersister represents the storage of the data
type TskDataPersister struct {
	// the data to persist
	Data TskData
}

func (p TskDataPersister) Read() {
}

func (p TskDataPersister) Write() {
	// marshal into json
	data, err := json.Marshal(p.Data)
	if err != nil {
		fmt.Errorf("Error on json.Marshal: %v", err)
		return
	}

	// write into file
	writeErr := ioutil.WriteFile(DefaultDataFileName, data, 0644)
	if err != nil {
		fmt.Errorf("Error on writing file: %v", writeErr)
		return
	}
}
