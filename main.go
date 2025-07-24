package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import for its side effects (registers sqlite3 driver)
)

var db *sql.DB

func main() {
	var err error
	databaseURL := "sqlite:///app/local.db" // Get this from docker-compose.yml
	db, err = sql.Open("sqlite3", databaseURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Ping the database to ensure connection is established
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Successfully connected to SQLite database!")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		email, err := getFirstUserEmail()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve user email",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!!",
			"first_user_email": email,
		})
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getFirstUserEmail() (string, error) {
	var email string
	row := db.QueryRow("SELECT email FROM users LIMIT 1")
	err := row.Scan(&email)
	if err == sql.ErrNoRows {
		return "No users found", nil
	}
	if err != nil {
		return "", err
	}
	return email, nil
}
