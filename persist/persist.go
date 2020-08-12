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
	Read() TskData
	Write(data TskData)
}

// TskDataPersister represents the storage of the data
type TskDataPersister struct{}

func (p TskDataPersister) Read() TskData {
	return TskData{}
}

func (p TskDataPersister) Write(data TskData) {
	// marshal into json
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("Error on json.Marshal: %v", err)
		return
	}

	// write into file
	writeErr := ioutil.WriteFile(DefaultDataFileName, jsonData, 0644)
	if err != nil {
		fmt.Errorf("Error on writing file: %v", writeErr)
		return
	}
}
