# tsk

[![version](https://img.shields.io/github/release/fitiavana07/tsk.svg?style=flat-square)](https://github.com/fitiavana07/tsk/releases/latest)

Task Management in the command line. Made with Go.

**Version 0.2 released: you can add tasks now**

[![asciicast](https://asciinema.org/a/354417.svg)](https://asciinema.org/a/354417)

## How it should work

### Scenario

```bash
$ tsk
No task, good news!

$ tsk add 'Implement task deletion'
Added: 1. Implement task deletion

$ tsk
Done:

Doing:

Todo:
    1. Implement task deletion

$ tsk do 1
Doing: 1. Implement task deletion

$ tsk
Done:

Doing:
    1. Implement task deletion

Todo:

$ tsk done 1
Done: 1. Implement task deletion

$ tsk
No task, good news!
```

### Details

Tasks are stored in a file named tasks.tsk in the current directory.

## Roadmap

- **v0.1**: task adding with tsk add and persistence
- **v0.2**: task list with tsk
- **v0.3**: add task states (do, doing, done)
- **v0.4**: advanced management (edit, delete)

## v0.1 Tasks Tracking

### Done

(Empty)

### Doing

- add new tasks in a file

### Todo

- [x] add license
- [x] number tasks
- [x] list tasks
- [ ] do a task
- [ ] mark task as done
- [ ] handle errors / wrong user input

### Notes

For directory structure, refer to https://github.com/golang-standards/project-layout.

### Ideas

Add progress bar using https://github.com/vbauerster/mpb.
Stylized output with https://github.com/Ullaakut/disgo.
Use Travis for CI.
Add colors using https://github.com/gookit/color
