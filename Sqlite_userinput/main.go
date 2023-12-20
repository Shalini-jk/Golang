package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open or create an SQLite database file
	db, err := sql.Open("sqlite3", "sample.db") // here we are creating the database inside the sqlite
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table if it doesn't exist
	// here we are creating the tables inside the above created database
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		email TEXT
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into the table (sql query to insert data)
	insertDataSQL := "INSERT INTO users (username, email) VALUES (?, ?)"

	// Take user input for a new user
	var newUsername, newEmail string // created variable to take input
	fmt.Print("Enter new username: ")
	fmt.Scan(&newUsername)
	fmt.Print("Enter new email: ")
	fmt.Scan(&newEmail)

	_, err = db.Exec(insertDataSQL, newUsername, newEmail)
	if err != nil {
		log.Fatal(err)
	}

	// Query data from the table
	queryDataSQL := "SELECT id, username, email FROM users"
	rows, err := db.Query(queryDataSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var username, email string
		err := rows.Scan(&id, &username, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", id, username, email)
	}
}
