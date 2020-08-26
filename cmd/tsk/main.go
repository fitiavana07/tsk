package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

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

	if len(args) == 1 {
		fmt.Fprintf(out, ui.MsgNoTask)
		return
	}

	if len(args) == 3 {
		taskName := args[2]
		task := fs.AddTask(taskName)
		ui.RenderTaskAdded(out, *task)
		return
	}
}
