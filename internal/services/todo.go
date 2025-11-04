package services

import (
	"database/sql"
	"fmt"
	"log"
	// "todo-golang/internal/models"
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
		// TODO: This can be done better with todo struct, then print array
		var id int
		var title string
		var done bool

		err := rows.Scan(&id, &title, &done)
		if err != nil {
		  fmt.Println("Failed to Executate: Could not obtain scan data correctly")
			log.Fatal(err)
			return
		}

		fmt.Printf("\n%d: %s\n   Done: %t\n", id, title, done)
	}

	if rowsEmpty {
		fmt.Println("No Tasks!")
		return
	}
}

func IsDoneTodo(db *sql.DB, done string) {
	// done will be boolean
	fmt.Println(done)
}

func DeleteTodo(db *sql.DB, id string) {
	// id will be int
	fmt.Println(id)
}

func UpdateTitleTodo(db *sql.DB, title string) {
	fmt.Println(title)
}

