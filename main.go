package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "todos.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	createQuery := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
	`
	_, err = db.Exec(createQuery)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("table created")

	_, err = db.Exec("INSERT INTO todos (name) VALUES (?)", "Learn Go")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("todo update")

	rows, err := db.Query("SELECT id, name FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Printf("Todo: %d - %s\n", id, name)
	}
}
