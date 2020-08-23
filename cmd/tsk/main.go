package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("not implemented")
}

// TskMain is the entrypoint after main()
func TskMain(args []string, stdout *bytes.Buffer) {
	fmt.Fprintf(stdout, welcomeMessage)
}
