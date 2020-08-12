// Main package
package main

import (
	//"encoding/json"
	"flag"
	"fmt"
	//"io/ioutil"
	"github.com/fitiavana07/tsk/cmd"
	//"github.com/fitiavana07/tsk/task"
)

// action names
const (
	ActionAdd  = "add"
	ActionDo   = "do"
	ActionDone = "done"
)

// main is the entrypoint in any go program
func main() {
	// TODO init file
	// tskDataFile.read()

	// parse command-line flags
	flag.Parse()

	// parse to param
	param := cmd.ParseArgs(flag.Args())
	fmt.Printf("%v\n", param)

	/*
		switch param.action {
		case "add":
			// adding a task
			// task := param.arg

			// add the task to the data
			// tskData := TskData{task}

			// data, err := json.Marshal(tskData)

			// if err != nil {
			// 	fmt.Printf("%v\n", err)
			// }

			// // overwrite the file
			// ioutil.WriteFile("tasks.tsk", data, 0644)

			// // tell the news
			// fmt.Printf("Added: %s\n", task)

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
	*/
}
