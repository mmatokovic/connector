package territory

// TerritoryService handles business logic for territories.
type TerritoryService struct{}

// NewTerritoryService creates a new instance of the service.
func NewTerritoryService() *TerritoryService {
	return &TerritoryService{}
}

// Returns the list of cities.
func (s *TerritoryService) listCities() []string {
	return []string{"Zagreb", "Split", "Rijeka", "Osijek"}
}

/*
// Returns the list of cities.
func listCities() {
	Repo.all(Cities)
}

// Gets a single inventory.
def get_inventory!(id), do: Repo.get!(Inventory, id)
func getCities(id) {
	Repo.get!(Inventory, id)
}

*/
