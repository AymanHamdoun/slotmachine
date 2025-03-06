package main

import (
	"log"
	"os"

	"go-backend/internal/app/api"
	"go-backend/internal/app/api/routes"
)

func main() {
	// Initialize the server
	server := api.NewServer()

	// Setup routes
	routes.SetupRoutes(server.Router())

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server starting on port %s", port)
	if err := server.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
