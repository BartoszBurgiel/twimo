package server

import (
	"net/http"
)

// Initialize the routes for a homepage of every location
// the route to the location is its ID
func (s *Server) initLocationHomepageRouter() {

	// Pull all locations from the database
	IDs, err := s.repo.GetLocationIDs()
	if err != nil {
		panic(err)
	}

	// Iterate over location names and create
	for _, id := range IDs {
		s.router.Route(Path("/location" + id))["GET"] = http.HandlerFunc(s.handleLocationHomepage)
	}
}
