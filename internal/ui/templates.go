package ui

const tmplTaskAdded = `Added:
  (use "tsk do {{ .ID }}" to start this task)
  (use "tsk" to show your Todo list)
        {{ .ID }}. {{ .Name }}
`

const tmplTasks = `Doing: None
Todo:
{{- range . }}
  (use "tsk do {{ .ID }}" to move into Doing)
        {{ .ID }}. {{ .Name }}
{{- end }}
`
