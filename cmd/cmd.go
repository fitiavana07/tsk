// Package cmd is the entrypoint
package cmd

import (
	//"flag"
	"fmt"
	"io"
)

// Main is the entrypoint of the program.
// Output is written into the provided io.Writer.
func Main(args []string, out io.Writer) {
	fmt.Fprint(out, "No task, good news!\n")
}
