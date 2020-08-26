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
