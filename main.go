package main

import (
	"database/sql"
	"log"

	db "tm/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

type apiConfig struct {
	DB *db.Queries
}

func main() {
	connection, err := sql.Open("sqlite3", "todos.db")

	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	apiConfig := apiConfig{
		DB: db.New(connection),
	}

	apiConfig.PrintTodos()
}
