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
  (Empty)

Doing:
  (Empty)

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

- **v0.1**: simple task management (add, do, done)

## v0.1 Tasks Tracking

### Done

(Empty)

### Doing

- add new tasks in a file

### Todo

- number tasks
- list tasks
- do a task
- mark task as done
