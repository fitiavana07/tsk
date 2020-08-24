package main

import (
	"bytes"
	"fmt"
	"os"
	// "path/filepath"

	"github.com/fitiavana07/tsk/internal/ui"
)

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
func TskMain(args []string, stdout *bytes.Buffer, file string) {

	rw, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Fprintf(stdout, ui.MsgWelcome)

	// TODO load data from rw, write "FileVersion": "1" to mark a first visit
}
