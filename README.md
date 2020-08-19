# tsk

Task Management in the command line. Made with Go.

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
  (Empty)

Doing:
  1. Implement task deletion

Todo:
  (Empty)

$ tsk done 1
Done: 1. Implement task deletion

$ tsk
Done:
  1. Implement task deletion

Doing:
  (Empty)

Todo:
  (Empty)
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

### Notes

For directory structure, refer to https://github.com/golang-standards/project-layout.
