// Package cmd is the entrypoint
package cmd

import (
	"fmt"
	"github.com/fitiavana07/tsk/internal/task"
	"io"
)

var tasksList = &[]task.Task{}

// Main is the entrypoint of the program.
// Output is written into the provided io.Writer.
// Get args from os.Args[1:] in command line
func Main(args []string, out io.Writer) {
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
	}
}
