// Package cmd is the entrypoint
package cmd

import (
	// "encoding/json"
	"fmt"
	"io"
	// "io/ioutil"
	// "os"
	// "path/filepath"

	"github.com/fitiavana07/tsk/internal/persist"
	"github.com/fitiavana07/tsk/internal/task"
)

var tasksList = &[]task.Task{}

// Main is the entrypoint of the program.
// Output is written into the provided io.Writer.
// Get args from os.Args[1:] in command line
func Main(
	out io.Writer,
	args []string,
	reader io.Reader,
	writer io.Writer,
) {
	switch l := len(args); {

	case l == 0:
		fmt.Fprintf(out, "No task, good news!\n")

	case l == 2:
		// current index
		index := len(*tasksList) + 1

		// add the task
		*tasksList = append(*tasksList, task.Task{index, args[1]})

		// output the response
		fmt.Fprintf(out, "Added: %d. %s\n", index, args[1])

		// TODO test it (added before telling "Added")
		// add to file
		tskData := &persist.TskData{}
		persist.ReadData(reader, tskData)
		fmt.Printf("present data: %+v\n", tskData)

		tskData.Tasks = append(tskData.Tasks, task.Task{index, args[1]})

		fmt.Printf("future data: %+v\n", tskData)
		persist.WriteData(tskData, writer)
	}
}
