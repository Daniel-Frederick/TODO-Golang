package main

import (
	"fmt"
	"os"
	"strings"
	"todo-golang/internal/database"
	"todo-golang/internal/services"
)

func main() {
	// Init sqlite
  db := database.InitDB()
	defer db.Close()
	
	// os.Args[0] = main.go
	// os.Args[1] = <cmd>
	// os.Args[2] = <cmd-optional> typically the todo ID 
	if 	len(os.Args) < 2 {
		fmt.Println("How to use: `go run ./cmd/todo <cmd> <cmd-optional>`")
		return
	}

	// TODO: don't use strings, make variables (objects or enums)
	cmd := Cmd(strings.ToLower(os.Args[1]))
	switch cmd {
	case CmdAdd:
		// TODO: Make sure to fix the strings in this file
		if len(os.Args) < 3 {
			fmt.Println("Did not specify new task!")
			return
		}
		services.AddTodo(db, os.Args[2])
		fmt.Printf("Added new task: %s\n", os.Args[2])
	case CmdList:
		fmt.Println("You entered the list cmd")
		services.ShowTodos(db)
	case CmdUpdate:
		fmt.Println("update cmd")
		if len(os.Args) < 3 {
			fmt.Println("Did not specify task ID!")
			return
		}
	case CmdDelete:
		fmt.Println("delete cmd")
		if len(os.Args) < 3 {
			fmt.Println("Did not specify task ID!")
			return
		}
	case CmdDone:
		fmt.Println("done cmd")
		if len(os.Args) < 3 {
			fmt.Println("Did not specify task ID!")
			return
		}
	default:
		fmt.Println("Not right: default hit!")
	}
}

