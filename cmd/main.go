package main

import (
	"fmt"
	"log"
	"menu360/config"
	"menu360/internal/user"
	"net/http"
)

func main() {
	fmt.Println("🚀 Server initializing...")

	// Load environment variables and connect to the database
	config.LoadEnv()
	db := config.ConnectDB()

	// ⚡ Initialize User module
	userModule := user.NewModule(db)

	// Configure router
	mux := http.NewServeMux()
	user.RegisterRoutes(mux, userModule.Controller) // Register routes for the User module

	// Start server
	port := ":8080"
	fmt.Printf("Server listening at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
