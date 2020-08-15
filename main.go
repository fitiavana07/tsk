package main

import (
	"os"

	"github.com/fitiavana07/tsk/cmd"
)

// main is the entrypoint in any go program
func main() {
	cmd.Main(os.Args[1:], os.Stdout)
}
