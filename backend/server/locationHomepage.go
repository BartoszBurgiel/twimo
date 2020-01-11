package server

import (
	"net/http"
	"strings"
)

// Initialize the routes for a homepage of every location
func (s *Server) initLocationHomepageRouter() {

	// Pull all locations from the database
	locationNames, err := s.repo.GetLocationNames()
	if err != nil {
		panic(err)
	}

	// Iterate over location names and create
	for _, name := range locationNames {
		s.router.Route(Path(processLocationsNameToRoute(name)))["GET"] = http.HandlerFunc(s.handleLocationHomepage)
	}
}

// Format locations name so that it can be used as a link
func processLocationsNameToRoute(name string) string {

	return "/" + strings.ReplaceAll(name, " ", "_")
}

func (s *Server) handleLocationHomepage(w http.ResponseWriter, r *http.Request) {
	/*
		Determine which location
	*/

	// Get the name from the url
	locationName := r.URL.String()

	// Add whitespace
	locationName = strings.ReplaceAll(locationName, "_", " ")

	// Delete the '/'
	locationName = locationName[1:]

	// Derive ID from the

}
