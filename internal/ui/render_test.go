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
	tsk := task.New(5, "do the 5th task")
	out := mockOut()

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
