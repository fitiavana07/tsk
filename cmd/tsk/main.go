package main

import (
	"bytes"
	"fmt"
	// "path/filepath"

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

func main() {
}

// TskMain is the entrypoint after main()
func TskMain(args []string, out *bytes.Buffer, file string) {
	fs := storage.NewFileStorage(file)
	if fs.IsFirstUse() {
		fmt.Fprintf(out, ui.MsgWelcome)
		fs.Save()
	} else {
		fmt.Fprintf(out, ui.MsgNoTask)
	}
}
