package services

import (
	"database/sql"
	"fmt"
	"log"
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

	for rows.Next() {
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

		fmt.Printf("%d: %s (done: %t)\n", id, title, done)
	}
}

