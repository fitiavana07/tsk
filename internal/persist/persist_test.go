package persist

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/fitiavana07/tsk/task"
)

// TODO handle errors: file not found on read, not json in file

func TestFilePersisterWrite(t *testing.T) {
	// data
	data := Data{
		LastTaskIndex: 1,
		Tasks: []task.Task{
			{
				Index: 1,
				Name:  "do this",
				State: task.StateDone,
			},
		},
	}

	// mock ioutil.WriteFile
	ioutilWriteFile := ioutil.WriteFile
	defer func() {
		ioutil.WriteFile = ioutilWriteFile
	}()
	fmt.Printf("%v\n", ioutilWriteFile)

	var persister Persister
	persister = FilePersister{"file.tsk"}

	// act
	persister.Write(data)

	// check
	// TODO check if the file was written properly
	// TODO check if the written data match the original data
}

func TestFilePersisterRead(t *testing.T) {
	var persister Persister
	persister = FilePersister{"tsk.json"}
	data := persister.Read()
	fmt.Printf("%v\n", data)
}
