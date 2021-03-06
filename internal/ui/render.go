package ui

import (
	"io"
	"text/template"

	"github.com/fitiavana07/tsk/pkg/task"
)

// RenderTaskAdded renders the messsage shown after successfully added a task.
func RenderTaskAdded(out io.Writer, t task.Task) {
	tmpl := template.Must(template.New("TaskAdded").Parse(tmplTaskAdded))
	tmpl.Execute(out, t)
}

// RenderTasks renders a list of tasks with their corresponding state.
func RenderTasks(out io.Writer, tasks []task.Task) {
	doings := filterTasksByState(tasks, task.StateDoing)
	todos := filterTasksByState(tasks, task.StateTodo)

	tmpl := template.Must(template.New("Tasks").Parse(tmplTasks))
	tmpl.Execute(out, map[string]([]task.Task){"doings": doings, "todos": todos})
}

// RenderDoneTasks renders a list of done tasks.
func RenderDoneTasks(out io.Writer, tasksDone []task.Task) {
	tmpl := template.Must(template.New("DoneTasks").Parse(tmplDoneTasks))
	tmpl.Execute(out, tasksDone)
}

// RenderTaskDoing renders the message to show just after moving a task into doing.
func RenderTaskDoing(out io.Writer, t task.Task) {
	tmpl := template.Must(template.New("TaskDoing").Parse(tmplTaskDoing))
	tmpl.Execute(out, t)
}

// RenderTaskDone renders the message to show just after moving a task into done.
func RenderTaskDone(out io.Writer, t task.Task) {
	tmpl := template.Must(template.New("TaskDone").Parse(tmplTaskDone))
	tmpl.Execute(out, t)
}

// TODO UNIT TEST
func filterTasksByState(tasks []task.Task, state task.State) (r []task.Task) {
	for _, t := range tasks {
		if t.State == state {
			r = append(r, t)
		}
	}
	return
}
