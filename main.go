package main

import (
	"database/sql"
	"log"
	"github.com/PrinceBabariya097/godb/internal/databse"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	connection, err := sql.Open("sqlite3", "todos.db")

	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	queries := 
}
