// Package cmd is the entrypoint
package cmd

import (
	"fmt"
	"io"
	// "strconv"

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

		fmt.Fprintf(out, "Done:\n")
		fmt.Fprintf(out, "\n")
		fmt.Fprintf(out, "Doing:\n")
		fmt.Fprintf(out, "\n")
		fmt.Fprintf(out, "Todo:\n")

		// list tasks
		for _, task := range tasks {
			fmt.Fprintf(out, "%5d. %s\n", task.Index, task.Name)
		}

	case l == 2:
		switch command := args[0]; command {
		case "add":
			// add to file
			index := tskData.LastTaskIndex + 1
			taskName := args[1]

			newTask := task.Task{
				Index: index,
				Name:  taskName,
			}

			// TODO refactor into tskData.addTodo(newTask)
			tskData.Tasks = append(tskData.Tasks, newTask)
			tskData.LastTaskIndex = index

			persist.WriteData(tskData, writer)

			fmt.Fprintf(out, "Added: %d. %s\n", index, taskName)

		case "do":
			// search
			// for _, t := range tskData.Tasks {
			// 	index, err := strconv.ParseInt(args[1], 10, 32)
			// 	if err != nil {
			// 		fmt.Printf("Error not yet handled: wrong arguments to tsk do")
			// 	}
			// 	if t.Index == int(index) {
			// 		t.State = task.StateDoing
			// 		break
			// 	}
			// }
		default:

		}
	}
}
