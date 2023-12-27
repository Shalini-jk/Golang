

//curl command

/* to get all users in json format:
 curl http://localhost:8080/users */

/*to accept data in SON format and post/ save it  into the database
curl -X POST -H "Content-Type: application/json" -d '{"username": "newuser", "email": "newuser@example.com"}' http://localhost:8080/users
*/

/*
 to delete data from database using curl command
 curl -X DELETE http://localhost:8080/users?id=1

*/

/* To update data from database using curl command
curl -X PUT -H "Content-Type: application/json" -d '{"username": "updateduser", "email": "updateduser@example.com"}' http://localhost:8080/users?id=2

*/


package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

// User represents a user's data
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	// Open or create an SQLite database file
	db, err := sql.Open("sqlite3", "sample.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table if it doesn't exist
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

	// Set up HTTP server
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUsersHandler(w, r, db)
		case http.MethodPost:
			addUserHandler(w, r, db)
		case http.MethodDelete:
			deleteUserHandler(w, r, db)
		case http.MethodPut:
			updateUserHandler(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// getUsersHandler retrieves all users from the database
func getUsersHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	queryDataSQL := "SELECT id, username, email FROM users"
	rows, err := db.Query(queryDataSQL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Retrieve user data
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Convert data to JSON
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Write JSON data to the response
	w.Write(jsonData)
}

// addUserHandler adds a new user to the database
func addUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Parse JSON request body
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert data into the table
	insertDataSQL := "INSERT INTO users (username, email) VALUES (?, ?)"
	_, err = db.Exec(insertDataSQL, newUser.Username, newUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User added successfully"))
}

// deleteUserHandler deletes a user from the database
func deleteUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the user ID from the request URL
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing user ID in the request", http.StatusBadRequest)
		return
	}

	// Convert userID to an integer
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Delete the user with the specified ID
	deleteDataSQL := "DELETE FROM users WHERE id = ?"
	_, err = db.Exec(deleteDataSQL, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

// updateUserHandler updates a user in the database
func updateUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the user ID from the request URL
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing user ID in the request", http.StatusBadRequest)
		return
	}

	// Convert userID to an integer
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Parse JSON request body
	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the user with the specified ID
	updateDataSQL := "UPDATE users SET username = ?, email = ? WHERE id = ?"
	_, err = db.Exec(updateDataSQL, updatedUser.Username, updatedUser.Email, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

