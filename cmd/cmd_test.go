package cmd

import (
	"testing"

	"bytes"
	"fmt"
	"strings"
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
	// test adding task given the args
	testArgs := func(args []string, t *testing.T) {
		// command line arguments, got from os.Args[1:]
		// an io.Writer implementation for tests
		buffer := &bytes.Buffer{}

		// act, in: args, out: buffer
		Main(args, buffer)

		// result
		got := buffer.String()

		// want
		want := fmt.Sprintf("Added: 1. %s\n", args[1])

		// check output
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	}

	argsCases := [][]string{
		{"add", "Implement task deletion"},
		{"add", "Write tests"},
	}
	for _, args := range argsCases {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			testArgs(args, t)
		})
	}
}
