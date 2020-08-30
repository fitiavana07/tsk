package ui

import (
	"testing"

	"bytes"

	"github.com/fitiavana07/tsk/pkg/task"
)

func mockOut() *bytes.Buffer {
	return &bytes.Buffer{}
}

func TestRenderTaskAdded(t *testing.T) {
	out := mockOut()
	tsk := task.New(5, "do the 5th task")

	RenderTaskAdded(out, *tsk)

	want := `Added:
  (use "tsk do 5" to start this task)
  (use "tsk" to show your Todo list)
        5. do the 5th task
`
	got := out.String()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}

}

func TestRenderTasks(t *testing.T) {
	out := mockOut()
	t.Run("SingleTodo", func(t *testing.T) {
		defer out.Reset()

		tasks := []task.Task{*task.New(7, "the seventh task")}

		RenderTasks(out, tasks)

		got := out.String()
		want := `Doing: None
Todo:
  (use "tsk do 7" to move into Doing)
        7. the seventh task
`
		if got != want {
			t.Errorf("wrong output:\n> got:\n%s\n> want:\n%s\n", got, want)
		}
	})

	t.Run("SingleDoing", func(t *testing.T) {
		defer out.Reset()

		tasks := []task.Task{
			task.Task{
				ID:    8,
				Name:  "The 8th task",
				State: task.StateDoing,
			},
		}

		RenderTasks(out, tasks)

		got := out.String()
		want := `Doing:
  (use "tsk done 8" to mark as Done)
        8. The 8th task
Todo: None
`
		if got != want {
			t.Errorf("wrong output:\n> got:\n%s\n> want:\n%s\n", got, want)
		}
	})
}

func TestRenderTaskDoing(t *testing.T) {
	out := mockOut()
	tsk := task.Task{
		ID:    3,
		Name:  "task to do",
		State: task.StateDoing,
	}

	RenderTaskDoing(out, tsk)

	want := `Moved into Doing:
  (once done, use "tsk done 3" to mark it as done)
        3. task to do
`
	got := out.String()
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s\n", got, want)
	}
}

func TestRenderTaskDone(t *testing.T) {
	out := mockOut()
	tsk := task.Task{
		ID:    5,
		Name:  "A done task",
		State: task.StateDone,
	}

	RenderTaskDone(out, tsk)

	want := `Marked as Done:
        5. A done task
Use "tsk --done" to show all done tasks
`

	got := out.String()
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s\n", got, want)
	}
}
