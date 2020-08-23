// Package persist contains all persistence utility functions of the program.
// Features will include:
// - encryption
package persist

import (
	"encoding/json"
	"github.com/fitiavana07/tsk/internal/task"
	"io"
	"io/ioutil"
)

// WriteData writes data into out.
func WriteData(data interface{}, out io.Writer) {
	// marshal into json
	jsonData, _ := json.Marshal(data)

	// write
	out.Write(jsonData)
}

// ReadData reads data from in
// Example:
//
// 		data := &struct{key string}{}
//		file, err := os.Open("file.json")
//		ReadData(file, data)
func ReadData(in io.Reader, data interface{}) {
	content, _ := ioutil.ReadAll(in)
	json.Unmarshal(content, data)
}

// TskData is the type of the whole data stored in persistence
type TskData struct {
	// list of tasks
	Tasks []task.Task

	// last task index, used to create the next task
	LastTaskIndex int
}

// TODO New, e.g: data.New(), or storage.NewData()
// TODO package or model
