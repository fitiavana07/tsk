package ui

const tmplTaskAdded = `Added:
  (use "tsk do {{ .ID }}" to start this task)
  (use "tsk" to show your Todo list)
        {{ .ID }}. {{ .Name }}
`
