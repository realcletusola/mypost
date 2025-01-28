package main 

import (
	"log"
	"net/http"

	"github.com/realcletusola/mypost/internal/db"
	"github.com/realcletusola/mypost/internal/routes"
)

func main() {
	// Initialize database connection 
	db.InitDB()
	defer db.DB.Close()

	// Run migrations 
	db.RunMigrations()

	// Setup and start the server 
	r := routes.SetupRouter()
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}