// Package cmd is the entrypoint
package cmd

import (
	//"flag"
	"fmt"
	"io"
)

// Main is the entrypoint of the program.
// Output is written into the provided io.Writer.
// Get args from os.Args[1:] in command line
func Main(args []string, out io.Writer) {
	switch l := len(args); {
	case l == 0:
		fmt.Fprint(out, "No task, good news!\n")
	case l == 2:
		fmt.Fprintf(out, "Added: 1. %s\n", args[1])
	}
}
