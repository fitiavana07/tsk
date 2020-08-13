package persist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fitiavana07/tsk/task"
)

// Data represents the whole tsk data
type Data struct {

	// last task index, used to create the next task
	LastTaskIndex int

	// list of tasks
	Tasks []task.Task
}

// Persister reads and writes data to external storage
type Persister interface {

	// Read reads data and returns a Data value
	Read() Data

	// Write writes data given a Data value
	Write(Data)
}

// FilePersister is a Persister which uses a single file as storage
type FilePersister struct {
	fileName string
}

// Read is FilePersister method to read data from file,
// it returns a data value
func (fp FilePersister) Read() (data Data) {
	fileContent, err := ioutil.ReadFile(fp.fileName)
	if err != nil {
		fmt.Errorf("Error on file read: %v", err)
	}
	json.Unmarshal(fileContent, &data)
	return
}

// Write is FilePersister method to write data into file,
func (fp FilePersister) Write(data Data) {

	// marshal into json
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("Error on json.Marshal: %v", err)
		return
	}

	// write into file
	writeErr := ioutil.WriteFile(fp.fileName, jsonData, 0644)
	if err != nil {
		fmt.Errorf("Error on writing file: %v", writeErr)
		return
	}
}
