package persist

import (
	"encoding/json"
	"fmt"
	"github.com/fitiavana07/tsk/internal/task"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// DefaultDataFilePath returns default data file path
// which is $HOME/.tsk/data.db
func DefaultDataFilePath() string {
	homeDir, _ := os.UserHomeDir()
	tskFile := filepath.Join(homeDir, ".tsk", "data.db")
	return tskFile
}

// WriteData writes data into out
func WriteData(data interface{}, out io.Writer) {
	// marshal into json
	jsonData, _ := json.Marshal(data)

	// write
	out.Write(jsonData)
}

// ReadData reads data from in
func ReadData(in io.Reader, data interface{}) {
	content, _ := ioutil.ReadAll(in)
	json.Unmarshal(content, data)
}

// FileReadWriter is implementation of io.ReadWriter to read and write to a file
// Example: WriteData(data, &FileWriter{"myFile.json"})
type FileReadWriter struct {
	filename string
}

func (frw FileReadWriter) Write(p []byte) (n int, err error) {
	err = ioutil.WriteFile(frw.filename, p, 0644)
	if err == nil {
		n = len(p)
	}
	return
}

func (frw FileReadWriter) Read(p []byte) (n int, err error) {
	content, err := ioutil.ReadFile(frw.filename)
	fmt.Printf("%s\n", content)
	fmt.Printf("%d\n", len(p))
	if err == nil {
		n = 0
		err = io.EOF
	}
	return
}

// TskData is the type of the whole data stored in persistence
type TskData struct {
	// list of tasks
	Tasks []task.Task
}

// Data represents the whole tsk data
// type Data struct {

// 	// last task index, used to create the next task
// 	LastTaskIndex int

// }
