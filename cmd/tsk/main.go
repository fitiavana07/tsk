package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/fitiavana07/tsk/internal/ui"
	"github.com/fitiavana07/tsk/pkg/storage"
)

// TODO used in real main()
// func defaultDataFilePath() string {
// 	homeDir, _ := os.UserHomeDir()
// 	tskFile := filepath.Join(homeDir, ".tsk", "data.db")
// 	return tskFile
// }
// func validateFile(filename string) {
// 	if _, err := os.Stat(filename); os.IsNotExist(err) {
// 		os.Mkdir(filepath.Dir(filename), os.ModePerm)
// 	}
// }

var datafile string

func init() {
	homeDir, _ := os.UserHomeDir()
	datafile = filepath.Join(homeDir, ".tskdata")
}

func main() {
	TskMain([]string{}, os.Stdout, datafile)
}

// TskMain is the entrypoint after main()
func TskMain(args []string, out io.Writer, file string) {
	fs := storage.NewFileStorage(file)
	if fs.IsFirstUse() {
		fmt.Fprintf(out, ui.MsgWelcome)
		fs.Save()
	} else {
		fmt.Fprintf(out, ui.MsgNoTask)
	}
}
