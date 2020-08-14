package tsk

import (
	"testing"

	"bytes"
)

func TestTskMain(t *testing.T) {
	buffer := &bytes.Buffer{}
	TskMain(buffer)

	got := buffer.String()
	want := "No task, good news!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
