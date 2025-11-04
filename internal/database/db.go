package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	_ "github.com/mattn/go-sqlite3"
	_ "embed"
)

// The `//go:embed` is a special comment
// sending the file schema.sql into the schema var
//go:embed schema.sql
var schema string

func InitDB() *sql.DB {
	// Create db directory
	dbDir := "data"
	dbPath := filepath.Join(dbDir, "todo.db")

	err := os.MkdirAll(dbDir, os.ModePerm)
	if err != nil {
		fmt.Println("failed to create new Directory")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("DB connection Failed")
		log.Fatal(err)
	}

	_, err = db.Exec(schema)
	if err != nil {
		fmt.Println("DB connection Failed")
		log.Fatal(err)
	}

	// fmt.Println("DB connection successful")
	return db
}

