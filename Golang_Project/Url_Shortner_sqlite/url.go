package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var shortURLDB *sql.DB

func main() {
	var input string
	var ch int
	var short_url string

	// Open SQLite database
	shortURLDB, _ = sql.Open("sqlite3", "./shorturl.db")
	defer shortURLDB.Close()

	// Create table if not exists
	createTable()

	for {
		fmt.Println("\n----URL Shortener Program----")
		fmt.Print("1. Get the short URL\n2. Expand the URL\n3. Print All URLs\n4. Exit\n")
		fmt.Print("Enter your choice to get the result: ")
		fmt.Scan(&ch)

		

		switch ch {
			
		case 1:
			fmt.Print("Enter the URL: ")
		    fmt.Scan(&input)
			short_url = shorten(input)
			insertURL(input, short_url)
			fmt.Println("Shorten URL is: ", short_url)
		case 2:
			fmt.Print("Enter the short URL: ")
			fmt.Scan(&short_url)
			original_url, found := expand(short_url)
			if found {
				fmt.Println("Expanded URL: ", original_url)
			} else { 
				fmt.Println("Short URL not found.")
			}
		case 3:
			printAllURLs()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func createTable() {
	_, err := shortURLDB.Exec(`CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		short_url TEXT UNIQUE,
		original_url TEXT
	);`)

	if err != nil {
		fmt.Println("Error creating table:", err)
		os.Exit(1)
	}
}
func shorten(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hasherValue := hasher.Sum(nil)
	hasherString := hex.EncodeToString(hasherValue)
	return hasherString[:12]
}

func expand(shortURL string) (string, bool) {
	var originalURL string
	err := shortURLDB.QueryRow("SELECT original_url FROM urls WHERE short_url = ?", shortURL).Scan(&originalURL)
	if err == sql.ErrNoRows {
		return "", false
	} else if err != nil {
		fmt.Println("Error expanding URL:", err)
		return "", false
	}
	return originalURL, true
}

func insertURL(originalURL, shortURL string) {
	_, err := shortURLDB.Exec("INSERT INTO urls (original_url, short_url) VALUES (?, ?)", originalURL, shortURL)
	if err != nil {
		fmt.Println("Error inserting URL:", err)
	}
}

func printAllURLs() {
	rows, err := shortURLDB.Query("SELECT short_url, original_url FROM urls")
	if err != nil {
		fmt.Println("Error fetching URLs:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\nAll URLs in the database:")
	for rows.Next() {
		var shortURL, originalURL string
		err := rows.Scan(&shortURL, &originalURL)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		fmt.Printf("Short URL: %s, Original URL: %s\n", shortURL, originalURL)
	}
}
