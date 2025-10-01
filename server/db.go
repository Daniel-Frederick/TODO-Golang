package database

import (
	"database/sql"
	"io/ioutil"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		log.Fatal(err)
	}

	todos, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err := db.Exec(string(todos))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
