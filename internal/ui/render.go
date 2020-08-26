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
	tmpl := template.Must(template.New("Tasks").Parse(tmplTasks))
	tmpl.Execute(out, tasks)
}
