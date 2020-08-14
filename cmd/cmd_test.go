package cmd

import (
	"testing"

	"bytes"
	"fmt"
)

// tests: tsk
func TestTskMain(t *testing.T) {
	// an io.Writer implementation for tests
	buffer := &bytes.Buffer{}

	// act
	Main([]string{}, buffer)

	// result
	got := buffer.String()

	// want
	want := "No task, good news!\n"

	// check
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// tests: tsk add 'Implement task deletion'
func TestTskMainAdd(t *testing.T) {
	// command line arguments, got from os.Args
	args := []string{"tsk", "add", "Implement task deletion"}

	// an io.Writer implementation for tests
	buffer := &bytes.Buffer{}

	// act, in: args, out: buffer
	Main(args, buffer)

	// result
	got := buffer.String()

	// want
	want := fmt.Sprintf("Added: 1. %s", args[2])

	// check output
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
