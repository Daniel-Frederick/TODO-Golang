package services

import (
	"database/sql"
	"fmt"
	"log"
	"todo-golang/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

// Add command
func AddTodo(db *sql.DB, task string) {
	insertSQL := "insert into todos (task, done) values (?, ?)"
	_, err := db.Exec(insertSQL, task, false)
	if err != nil {
		fmt.Println("Failed to Executate: nothing added")
		fmt.Println(err)
		return
	}
	fmt.Println("Add TODO successful")
}

// List command
func ShowTodos(db *sql.DB) {
	showTodos := "select * from todos;"
	rows, err := db.Query(showTodos)
	if err != nil {
		fmt.Println("Failed to Executate: Could not obtain data")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	rowsEmpty := true

	for rows.Next() {
		rowsEmpty = false
		var todo models.Todo

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Done)
		if err != nil {
		  fmt.Println("Failed to Executate: Could not obtain scan data correctly")
			log.Fatal(err)
			return
		}

		fmt.Printf("\n%d: %s\n   Done: %t\n", todo.Id, todo.Title, todo.Done)
	}

	if rowsEmpty {
		fmt.Println("No Tasks!")
		return
	}
}

func IsDoneTodo(db *sql.DB, id int) {
	fmt.Println("id: ", id)

	// Query for the specific row with id, not all rows
	showTodos := "select id, done from todos where id = ?;"
	todo, err := db.Query(showTodos, id)
	if err != nil {
	  fmt.Println("Failed to Executate: Could not obtain scan data correctly")
		log.Fatal(err)
		return
	}

	updateDoneSQL := "update todos set done = ? where id = ?"
	_, err := db.Exec(updateDoneSQL, !todo.Done, todo.Id)

	fmt.Printf("\n%d: Done: %t\n", todo.Id, todo.Done)
}

func DeleteTodo(db *sql.DB, id string) {
	fmt.Println(id)
}

func UpdateTitleTodo(db *sql.DB, id int, title string) {
	fmt.Println("title: ", title, "\nid: ", id)
}

func HelpTodo() {
	fmt.Println("Available CMDs: go run ./cmd/todo")
	fmt.Println("  help")
	fmt.Println("  list")
	fmt.Println("  add <task>")
	fmt.Println("  update <id> <new task>")
}

