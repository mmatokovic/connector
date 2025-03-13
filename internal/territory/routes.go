package territory

import "net/http"

// RegisterRoutes registers the HTTP routes for the territory module.
func RegisterRoutes(handler *TerritoryHandler) {
	http.HandleFunc("/cities", handler.indexCities)
	//http.HandleFunc("/cities/", handler.showCity)
}
