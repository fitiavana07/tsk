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

// Main is the entrypoint of the program.
// Output is written into the provided io.Writer.
// Get args from os.Args[1:] in command line
func Main(
	out io.Writer,
	args []string,
	reader io.Reader,
	writer io.Writer,
) {
	tskData := &persist.TskData{}
	persist.ReadData(reader, tskData)

	switch l := len(args); {

	case l == 0:
		tasks := tskData.Tasks
		if len(tasks) == 0 {
			fmt.Fprintf(out, "No task, good news!\n")
			return
		}

		fmt.Fprintf(out, "Todo:\n")
		// list tasks
		for _, task := range tasks {
			fmt.Fprintf(out, "  %d. %s\n", task.Index, task.Name)
		}

	case l == 2:
		// add to file
		index := tskData.LastTaskIndex + 1
		taskName := args[1]

		newTask := &task.Task{
			Index: index,
			Name:  taskName,
		}

		tskData.Tasks = append(tskData.Tasks, *newTask)
		tskData.LastTaskIndex = index

		persist.WriteData(tskData, writer)

		fmt.Fprintf(out, "Added: %d. %s\n", index, taskName)

	}
}
