package ui

const tmplTaskAdded = `Added:
  (use "tsk do {{ .ID }}" to start this task)
  (use "tsk" to show your Todo list)
        {{ .ID }}. {{ .Name }}
`

const tmplTasks = `
{{- if .doings -}}
Doing:

{{- range .doings }}
  (use "tsk done {{ .ID }}" to mark as Done)
        {{ .ID }}. {{ .Name }}
{{ end -}}

{{- else -}}
Doing: None
{{ end -}}

{{- if .todos -}}
Todo:

{{- range .todos }}
  (use "tsk do {{ .ID }}" to move into Doing)
        {{ .ID }}. {{ .Name }}
{{ end -}}

{{- else -}}
Todo: None
{{ end -}}
`

const tmplTaskDoing = `Moved into Doing:
  (once done, use "tsk done {{ .ID }}" to mark it as done)
        {{ .ID }}. {{ .Name }}
`
