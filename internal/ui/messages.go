package ui

const (
	// MsgWelcome is the message shown on the first use of the app
	MsgWelcome = `Hi!
Thanks for using tsk!
To get started, add a task using: tsk add "your task".
`

	// MsgNoTask is the message shown when the user wants the list of his tasks
	// with `tsk` but there is currently no task in todo, nor in doing
	MsgNoTask = `Your todo list is empty
To add a task, use: tsk add "your task"
`

	// MsgAllDone is the message shown when all tasks are done.
	MsgAllDone = `Congrats! All done for today.
Use "tsk --done" to show all done tasks
`
)
