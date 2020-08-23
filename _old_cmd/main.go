package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fitiavana07/tsk/cmd"
)

// DefaultDataFilePath returns default data file path
// which is $HOME/.tsk/data.db
func defaultDataFilePath() string {
	homeDir, _ := os.UserHomeDir()
	tskFile := filepath.Join(homeDir, ".tsk", "data.db")
	return tskFile
}

func validateFile(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Mkdir(filepath.Dir(filename), os.ModePerm)
	}

}

// main is the entrypoint in any go program
func main2() {
	filePath := defaultDataFilePath()
	validateFile(filePath)

	reader, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	writer, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	cmd.Main(os.Stdout, os.Args[1:], reader, writer)
}
