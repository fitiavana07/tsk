package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fitiavana07/tsk/internal/ui"
	"github.com/fitiavana07/tsk/pkg/storage"
)

func init() {
}

func main() {
	homeDir, _ := os.UserHomeDir()
	datafile := filepath.Join(homeDir, ".tskdata")
	TskMain(os.Args, os.Stdout, datafile)
}

// TskMain is the entrypoint after main()
func TskMain(args []string, out io.Writer, file string) {
	fs := storage.NewFileStorage(file)

	if fs.IsFirstUse() {
		fmt.Fprintf(out, ui.MsgWelcome)
		fs.Save()
		return
	}

	// tsk
	if len(args) == 1 {
		tasks := fs.Tasks()
		if len(tasks) == 0 {
			tasksDone := fs.TasksDone()
			if len(tasksDone) == 0 {
				fmt.Fprintf(out, ui.MsgNoTask)
			} else {
				fmt.Fprintf(out, ui.MsgAllDone)
			}
		} else {
			ui.RenderTasks(out, tasks)
		}
		return
	}

	// tsk --done
	if len(args) == 2 {
		doneTasks := fs.TasksDone()
		ui.RenderDoneTasks(out, doneTasks)
	}

	// tsk cmd arg
	if len(args) == 3 {
		cmd := args[1]
		arg := args[2]

		switch cmd {

		case "add":
			task := fs.AddTask(arg)
			ui.RenderTaskAdded(out, *task)
			return

		case "do":
			taskID, _ := strconv.ParseInt(arg, 10, 0)
			task := fs.DoTask(int(taskID))
			ui.RenderTaskDoing(out, *task)
			return

		case "done":
			taskID, _ := strconv.ParseInt(arg, 10, 0)
			task := fs.DoneTask(int(taskID))
			ui.RenderTaskDone(out, *task)
		default:
			// TODO handle unrecognized command

		}

		return
	}
}
