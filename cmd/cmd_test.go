package cmd

import (
	"testing"

	"bytes"
	"fmt"
	// "os"
	"github.com/fitiavana07/tsk/internal/persist"
	"github.com/fitiavana07/tsk/internal/task"
)

// TestTskMainFirstUsage tests first use of the program
func TestTskMainFirstUsage(t *testing.T) {
	// an io.Writer implementation for tests
	buffer := &bytes.Buffer{}
	readerBuffer := &bytes.Buffer{}
	writerBuffer := &bytes.Buffer{}

	// Given: There is no data file

	// When:
	Main(buffer, []string{}, readerBuffer, writerBuffer)

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
func TestTskMainAddAfterNoDataFile(t *testing.T) {
	// Given: the file doesn't exist
	persistBuffer := &bytes.Buffer{}

	// test adding task given the args
	testArgs := func(args []string, index int, t *testing.T) {
		// Given: args: command line arguments as {name}

		// buffer to write output
		printBuffer := &bytes.Buffer{}

		// When: I run tsk with args
		Main(printBuffer, args, persistBuffer, persistBuffer)

		// result
		got := printBuffer.String()

		// Then: output: Added: {index}. {name}
		want := fmt.Sprintf("Added: %d. %s\n", index, args[1])

		// check output
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
		// test for the cases
		for index, task := range tasks {
			testArgs([]string{"add", task}, index+1, t)
		}
	})

	t.Run("TasksAreInFile", func(t *testing.T) {
		// read file
		data := &persist.TskData{}
		persist.ReadData(persistBuffer, data)

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

func TestTskMainAddPersistentNonEmpty(*testing.T) {
	// 1. given a json data file with a non-empty tasks list,
	// 		where the last index is 5
	// 2. given an args from command line: ["add", "Next task"]
	// 3. verify the output is: "Added: 6. Next task"
	// 4. verify the task is present in the json data file
}

func TestTskMainListPresentTodoTasks(t *testing.T) {
	// given a non-empty data
	data := &persist.TskData{
		LastTaskIndex: 2,
		Tasks: []task.Task{
			{Index: 1, Name: "add 2 tasks"},
			{Index: 2, Name: "list tasks"},
		},
	}
	persistBuffer := &bytes.Buffer{}
	persist.WriteData(data, persistBuffer)

	// buffer to write output
	printBuffer := &bytes.Buffer{}

	// When: I run tsk with the data
	Main(printBuffer, []string{}, persistBuffer, persistBuffer)

	want := `Todo:
    1. add 2 tasks
    2. list tasks
`
	got := printBuffer.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
