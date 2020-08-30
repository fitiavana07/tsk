package ui

const tmplTaskAdded = `Added:
  (use "tsk do {{ .ID }}" to start this task)
  (use "tsk" to show your Todo list)
        {{ .ID }}. {{ .Name }}
`

const tmplTasks = `
{{- if .doings -}}
Doing:{{ "\n" }}

	{{- if len .doings | eq 1 -}}

		{{- with index .doings 0 -}}
			{{- "  " }}(use "tsk done {{ .ID }}" to mark as Done){{ "\n" }}
			{{- "        "}}{{ .ID }}. {{ .Name }}{{ "\n" }}
		{{- end -}}

	{{- else -}}
		{{- "  " }}(use "tsk done <n>" to mark as Done){{ "\n" }}

		{{- range .doings -}}
			{{- "        " }}{{ .ID }}. {{ .Name }}{{ "\n" }}
		{{- end -}}

	{{- end -}}

{{- else -}}
Doing: None{{ "\n" }}
{{- end -}}

{{- if .todos -}}
Todo:{{ "\n" }}

	{{- if len .todos | eq 1 -}}

		{{- with index .todos 0 -}}
			{{- "  " }}(use "tsk do {{ .ID }}" to move into Doing){{ "\n" }}
			{{- "        " }}{{ .ID }}. {{ .Name }}{{ "\n" }}
		{{- end -}}

	{{- else -}}
		{{- "  " }}(use "tsk do <n>" to move into Doing){{ "\n" }}

		{{- range .todos -}}
			{{- "        " }}{{ .ID }}. {{ .Name }}{{ "\n" }}
		{{- end -}}

	{{- end -}}

{{- else -}}
Todo: None{{ "\n" }}
{{- end -}}
`

const tmplTaskDoing = `Moved into Doing:
  (once done, use "tsk done {{ .ID }}" to mark it as done)
        {{ .ID }}. {{ .Name }}
`
