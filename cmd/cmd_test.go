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
	Main(buffer, []string{})

	// result
	got := buffer.String()

	// want
	want := "No task, good news!\n"

	// check
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestTskMainAdd tests adding a task without pre-added tasks
func TestTskMainAdd(t *testing.T) {
	// test adding task given the args
	testArgs := func(args []string, index int, t *testing.T) {
		// Given: args: command line arguments as {name}

		// buffer to write output
		buffer := &bytes.Buffer{}

		// When: I run tsk with args
		Main(buffer, args)

		// result
		got := buffer.String()

		// Then: output: Added: {index}. {name}
		want := fmt.Sprintf("Added: %d. %s\n", index, args[1])

		// check output
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	// multiple test cases to verify that output depend on intput
	argsCases := [][]string{
		{"add", "Implement task deletion"},
		{"add", "Write tests"},
	}

	// test for the cases
	for index, args := range argsCases {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			testArgs(args, index+1, t)
		})
	}

	// TODO verify all tasks are in the data file
}

func TestTskMainAddPersistentNonEmpty(*testing.T) {
	// 1. given a json data file with a non-empty tasks list,
	// 		where the last index is 5
	// 2. given an args from command line: ["add", "Next task"]
	// 3. verify the output is: "Added: 6. Next task"
	// 4. verify the task is present in the json data file
}
