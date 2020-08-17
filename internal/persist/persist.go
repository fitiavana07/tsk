package persist

import (
	"encoding/json"
	//"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	//"github.com/fitiavana07/tsk/internal/task"
)

// DefaultDataFilePath returns default data file path
// which is $HOME/.tsk/data.db
func DefaultDataFilePath() string {
	homeDir, _ := os.UserHomeDir()
	tskFile := filepath.Join(homeDir, ".tsk", "data.db")
	return tskFile
}

// WriteData writes data into file filename
func WriteData(data Data, out io.Writer) {
	// marshal into json
	jsonData, _ := json.Marshal(data)

	// write
	out.Write(jsonData)
}

// FileWriter is implementation of io.Writer to write to a file
// Example: WriteData(data, &FileWriter{"myFile.json"})
type FileWriter struct {
	filename string
}

func (fw FileWriter) Write(p []byte) (n int, err error) {
	err = ioutil.WriteFile(fw.filename, p, 0644)
	if err != nil {
		n = len(p)
	}
	return
}

// Data is any data that can be written into json
type Data interface{}

// Data represents the whole tsk data
// type Data struct {

// 	// last task index, used to create the next task
// 	LastTaskIndex int

// 	// list of tasks
// 	Tasks []task.Task
// }
