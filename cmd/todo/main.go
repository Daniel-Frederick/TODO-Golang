package main

import (
	"fmt"
	"os"
	"todo-golang/internal/database"
	"todo-golang/internal/services"
)

func main() {
	// Init sqlite
  db := database.InitDB()
	defer db.Close()
	
	// os.Args[0] = main.go
	// os.Args[1] = <task> or <id>
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
		services.AddTodo(db, os.Args[2])
		fmt.Printf("Added new task: %s\n", os.Args[2])
	case "list":
		fmt.Println("You entered the list cmd")
		services.ShowTodos(db)
	
	// TODO: Continue creating extending switch to add more cmds
	}
}

