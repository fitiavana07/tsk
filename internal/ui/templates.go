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
			{{- template "TipToDone" .ID -}}
			{{- template "Item" . -}}
		{{- end -}}

	{{- else -}}
		{{- template "TipToDone" "<n>" -}}

		{{- range .doings -}}
			{{- template "Item" . -}}
		{{- end -}}
	{{- end -}}

{{- else -}}
Doing: None{{ "\n" }}
{{- end -}}

{{- if .todos -}}
Todo:{{ "\n" }}

	{{- if len .todos | eq 1 -}}

		{{- with index .todos 0 -}}
			{{- template "TipToDoing" .ID -}}
			{{- template "Item" . -}}
		{{- end -}}

	{{- else -}}
		{{- template "TipToDoing" "<n>" -}}

		{{- range .todos -}}
			{{- template "Item" . -}}
		{{- end -}}
	{{- end -}}

{{- else -}}
Todo: None{{ "\n" }}
{{- end -}}

{{- define "Item" -}}
	{{- "        "}}{{ .ID }}. {{ .Name }}{{ "\n" }}
{{- end -}}

{{- define "TipToDone" -}}
	{{- "  " }}(use "tsk done {{ . }}" to mark as Done){{ "\n" }}
{{- end -}}

{{- define "TipToDoing" -}}
	{{- "  " }}(use "tsk do {{ . }}" to move into Doing){{ "\n" }}
{{- end -}}
`

const tmplTaskDoing = `Moved into Doing:
  (once done, use "tsk done {{ .ID }}" to mark it as done)
        {{ .ID }}. {{ .Name }}
`

const tmplTaskDone = `Marked as Done:
        {{ .ID }}. {{ .Name }}
Use "tsk --done" to show all done tasks
`
