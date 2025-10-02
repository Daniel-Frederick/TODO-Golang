package database

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		fmt.Println("DB connection Failed")
		log.Fatal(err)
	}

	todos, err := os.ReadFile("database/schema.sql")
	if err != nil {
		fmt.Println("DB connection Failed")
		log.Fatal(err)
	}

	_, err = db.Exec(string(todos))
	if err != nil {
		fmt.Println("DB connection Failed")
		log.Fatal(err)
	}

	fmt.Println("DB connection successful")
	return db
}

func Test() {
	fmt.Println("TESTING DATABASE")
}
