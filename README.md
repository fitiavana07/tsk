# tsk

Task Management in the command line. Made with Go.

## How it should work

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

$ tsk
Done: 1. Implement task deletion

$ tsk
Done:
  1. Implement task deletion

Doing:
  (Empty)

Todo:
  (Empty)
```

