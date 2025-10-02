package services

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func AddTodo(db *sql.DB, task string) {
	fmt.Println("AddTODO")
	_, err := db.Exec("insert into todos (task, done) values (?, ?)", task, false)
	if err != nil {
		fmt.Println("Failed to Executate: nothing added")
		fmt.Println(err)
		return
	}
	fmt.Println("Add TODO successful")

	// RowsAffected, _ = insert.RowsAffected()
	// fmt.Println("RowsAffected: " + RowsAffected)
}

