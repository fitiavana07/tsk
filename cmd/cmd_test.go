package cmd

import (
	"testing"

	"bytes"
	"fmt"

	"github.com/fitiavana07/tsk/internal/persist"
	"github.com/fitiavana07/tsk/internal/task"
)

// mock stdout
func mockStdout() *bytes.Buffer {
	return &bytes.Buffer{}
}

// mock file reader/writer
func mockFileRW() *bytes.Buffer {
	return &bytes.Buffer{}
}

// TestTskMainFirstUsage tests first use of the program
func TestTskMainFirstUsage(t *testing.T) {
	stdout := mockStdout()

	// Given: There is no data file
	fileRW := mockFileRW()

	// When:
	Main(stdout, []string{}, fileRW, fileRW)

	// Then:
	got := stdout.String()
	want := "No task, good news!\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestTskMainAdd tests adding a task without pre-added tasks
func TestTskMainAddAfterNoDataFile(t *testing.T) {
	stdout := mockStdout()

	// Given: the file doesn't exist
	fileRW := mockFileRW()

	// test adding task given the args
	testArgs := func(args []string, index int, t *testing.T) {
		// clean stdout after each command
		defer stdout.Reset()

		// Given: args: command line arguments as {name}
		// When: I run tsk with args
		Main(stdout, args, fileRW, fileRW)

		// Then: output: Added: {index}. {name}
		got := stdout.String()
		want := fmt.Sprintf("Added: %d. %s\n", index, args[1])
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	}

	// multiple test cases to verify that output depend on intput
	tasks := []string{
		"Implement task deletion",
		"Write tests",
	}

	t.Run("PrintCorrectOutput", func(t *testing.T) {
		for index, task := range tasks {
			testArgs([]string{"add", task}, index+1, t)
		}
	})

	t.Run("TasksAreInFile", func(t *testing.T) {
		// read data
		data := &persist.TskData{}
		persist.ReadData(fileRW, data)

		// find the tasks
		for _, task := range tasks {
			found := false
			for i := range data.Tasks {
				if data.Tasks[i].Name == task {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("task %q not found in data file", task)
			}
		}
	})

}

// TestTskMainListPresentTodoTasks tests listing present tasks after adding some tasks
func TestTskMainListPresentTodoTasks(t *testing.T) {
	stdout := mockStdout()
	fileRW := mockFileRW()

	// given a non-empty data
	data := &persist.TskData{
		LastTaskIndex: 2,
		Tasks: []task.Task{
			{Index: 1, Name: "add 2 tasks"},
			{Index: 2, Name: "list tasks"},
		},
	}
	persist.WriteData(data, fileRW)

	// When: I run tsk with the data
	Main(stdout, []string{}, fileRW, fileRW)

	want := `Done:

Doing:

Todo:
    1. add 2 tasks
    2. list tasks
`
	got := stdout.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
