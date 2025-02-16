package main

import (
	"log"
	"menu360/config"
	"menu360/internal/user"
	"net/http"
)

func main() {
	// Load environment variables and connect to the database
	config.LoadEnv()
	config.ConnectDB()

	// Create the HTTP server
	mux := http.NewServeMux()

	// Initialize user module and register routes
	userModule := user.NewModule()
	userModule.RegisterRoutes(mux)

	// Start server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
