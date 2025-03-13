package main

import (
	"connector/internal/territory"
	"log"
	"net/http"
)

func main() {
	// Initialize service and handler
	territoryService := territory.NewTerritoryService()
	territoryHandler := territory.NewTerritoryHandler(territoryService)

	// Register routes
	territory.RegisterRoutes(territoryHandler)

	// Start the HTTP server
	port := ":8080"
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Defines connector scope
//fun router
