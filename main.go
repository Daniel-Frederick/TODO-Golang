package main

import (
	"fmt"
	"os"
	"todo-golang/database"
	"todo-golang/services"
)

func main() {
	// Init sqlite
	// db := initDB()
	
	// os.Args[0] = main.go
	// os.Args[1] = <task> or <id>
	fmt.Println("TODO App started!")
	if 	len(os.Args) < 2 {
		fmt.Println("How to use: idk")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Did not specify new task!")
			return
		}
		fmt.Println("Added new task: %s", os.Args[3])
	case "list":
		fmt.Println("You entered the list cmd")
		// TODO: Connect to database to return all tasks
	
	// TODO: Continue creating extending switch to add more cmds

	}

}

