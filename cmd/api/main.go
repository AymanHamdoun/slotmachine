package main

import (
	"go-backend/internal/app/api"
	"go-backend/internal/app/api/routes"
	"go-backend/internal/app/config"
	"log"
)

func main() {
	// Initialize the server
	server := api.NewServer()
	config.Load()

	// Setup routes
	routes.SetupRoutes(server.Router())

	// Start the server
	log.Printf("Server starting on port %s", config.Get().AppPort)
	if err := server.Start(":" + config.Get().AppPort); err != nil {
		log.Fatal(err)
	}
}
