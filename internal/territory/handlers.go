package territory

import (
	"encoding/json"
	"net/http"
)

// TerritoryHandler handles HTTP requests for territories.
type TerritoryHandler struct {
	service *TerritoryService
}

// NewTerritoryHandler creates a new handler instance.
func NewTerritoryHandler(service *TerritoryService) *TerritoryHandler {
	return &TerritoryHandler{service: service}
}

// HandleCities responds to GET /territory/cities
func (h *TerritoryHandler) indexCities(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	response, _ := json.Marshal(h.service.listCities())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

/*
func showCity(conn, %{"id" => id}) {
	inventory = get_inventory!(id)
	render(conn, :show, inventory: inventory)
}
*/
