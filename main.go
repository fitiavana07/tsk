// Main file
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

var hasTask = true

var todos []string

type TskData struct {
	Todo string
}

type CLIOption struct {
	action string
	arg    string
}

// main is the entrypoint in any go program
func main() {
	if !hasTask {
		fmt.Println("No task, good news!")
	}

	flag.Parse()

	action := flag.Arg(0)
	args := flag.Args()
	fmt.Printf("%v\n", args)

	switch action {
	case "add":
		// adding a task
		task := flag.Arg(1)

		// add the task to the data

		tskData := TskData{task}

		data, err := json.Marshal(tskData)

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		// overwrite the file
		ioutil.WriteFile("tasks.tsk", data, 0644)

		// tell the news
		fmt.Printf("Added: %s\n", task)

	case "do":
		// doing a task
		fmt.Println("You choose to do")

	case "done":
		// done a task
		fmt.Println("You choose to done")

	default:
		// list
		fmt.Println("You choose to list")
	}
}
