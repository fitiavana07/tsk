package main

import (
	"testing"

	"bytes"
)

/*
Functional testing of the tsk application.
These tests should be compiled in one func, but divided into subtests

$ tsk
Hi!
Thanks for using tsk!
To get started, add a task using: tsk add "your task".
$ tsk
No task yet
To add a task, use: tsk add "your task"
$ tsk add "Call John Doe"
Added:
  (use "tsk do 1" to start this task)
  (use "tsk" to show your Todo list)
        1. Call John Doe
$ tsk
Doing: None
Todo:
  (use "tsk do 1" to move into Doing)
        1. Call John Doe
$ tsk do 1
Moved into Doing:
  (once done, use "tsk done 1" to mark it as done)
        1. Call John Doe
$ tsk
Doing:
  (use "tsk done 1" to mark as Done)
		1. Call John Doe
Todo: None
$ tsk add "Create a PR"
Added:
  (use "tsk do 2" to start this task)
  (use "tsk" to show your Todo list)
        2. Create a PR
$ tsk
Doing:
  (use "tsk done 1" to mark as Done)
		1. Call John Doe
Todo:
  (use "tsk do 2" to move into Doing)
        2. Create a PR
$ tsk do 2
Moved into Doing:
  (use "tsk done 2" to mark as Done)
        2. Create a PR
$ tsk
Doing:
  (use "tsk done <n>" to mark as Done)
        1. Call John Doe
        2. Create a PR
Todo: None
$ tsk done 2
Marked as Done:
        2. Create a PR
Use "tsk --done" to show all done tasks
$ tsk
Doing:
  (use "tsk done 1" to mark as Done)
        1. Call John Doe
Todo: None
$ tsk done 1
Marked as Done:
        1. Call John Doe
Use "tsk --done" to show all done tasks
$ tsk
Congrats! All done for today.
Use "tsk --done" to show all done tasks
$ tsk --done
Done:
	1. Call John Doe
	2. Create a PR
*/

var scenario = []act{{
	"FirstUse",
	a("tsk"),
	`Hi!
Thanks for using tsk!
To get started, add a task using: tsk add "your task".
`}, {
	"SecondUse",
	a("tsk"),
	`No task yet
To add a task, use: tsk add "your task"
`}, {
	"AddFirstTask",
	a("tsk", "add", "Call John Doe"),
	`Added:
  (use "tsk do 1" to start this task)
  (use "tsk" to show your Todo list)
        1. Call John Doe
`}, {
	"ListTaskAfterAdd",
	a("tsk"),
	`Doing: None
Todo:
  (use "tsk do 1" to move into Doing)
        1. Call John Doe
`}, {
	"DoFirstTask",
	a("tsk", "do", "1"),
	`Moved into Doing:
  (once done, use "tsk done 1" to mark it as done)
        1. Call John Doe
`}, {
	"ListTaskAfterDo",
	a("tsk"),
	`Doing:
  (use "tsk done 1" to mark as Done)
		1. Call John Doe
Todo: None
`}, {
	"AddSecondTask",
	a("tsk", "add", "Create a PR"),
	`Added:
  (use "tsk do 2" to start this task)
  (use "tsk" to show your Todo list)
        2. Create a PR
`}, {
	"ListTaskAfterAddSecond",
	a("tsk"),
	`Doing:
  (use "tsk done 1" to mark as Done)
		1. Call John Doe
Todo:
  (use "tsk do 2" to move into Doing)
        2. Create a PR
`}, {
	"DoSecondTask",
	a("tsk", "do", "2"),
	`Moved into Doing:
  (use "tsk done 2" to mark as Done)
        2. Create a PR
`}, {
	"ListTaskAfterDoSecondTask",
	a("tsk"),
	`Doing:
  (use "tsk done <n>" to mark as Done)
        1. Call John Doe
        2. Create a PR
Todo: None
`}, {
	"DoneSecondTask",
	a("tsk", "done", "2"),
	`Marked as Done:
        2. Create a PR
Use "tsk --done" to show all done tasks
`}, {
	"ListTaskAfterDoneSecondTask",
	a("tsk"),
	`Doing:
  (use "tsk done 1" to mark as Done)
        1. Call John Doe
Todo: None
`}, {
	"DoneFirstTask",
	a("tsk", "done", "1"),
	`Marked as Done:
        1. Call John Doe
Use "tsk --done" to show all done tasks
`}, {
	"ListTaskAfterDoneFirstTask",
	a("tsk"),
	`Congrats! All done for today.
Use "tsk --done" to show all done tasks
`}, {
	"ListDoneTasks",
	a("tsk", "--done"),
	`Done:
	1. Call John Doe
	2. Create a PR
`}}

func TestFunctional(t *testing.T) {
	stdout := mockStdout()
	for _, ac := range scenario {
		t.Run(ac.name, func(t *testing.T) {
			testAct(t, ac, stdout)
		})
	}
}

func testAct(t *testing.T, ac act, stdout *bytes.Buffer) {
	defer stdout.Reset()

	TskMain(ac.args, stdout)

	got := stdout.String()
	want := ac.output

	if got != want {
		t.Errorf("wrong output: %q, want %q", got, want)
	}
}

func a(args ...string) []string {
	return args
}

type act struct {
	name   string
	args   []string
	output string
}

func mockStdout() *bytes.Buffer {
	return &bytes.Buffer{}
}
