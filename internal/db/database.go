package db 

import (
	"database/sql"
	"fmt"
	"log"

	_"github.com/lib/pq" // PostgreSQL driver 
)

var DB *sql.DB 

// Initialize PostgreSQL database connection
func InitDB() {
	connStr := "postgres://postgres:postgres@localhost:5432/mypost?sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Verify the connection 
	err := DB.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	log.Println("Database connection established.")
}

// Run database migrations 
func RunMigrations() {
	query := `
		CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			content TEXT NOT NULL,
			date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			category_id INT REFERENCES categories(id) ON DELETE SET NULL
		);

	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migrations applied successfully.")
}


