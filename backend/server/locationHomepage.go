package server

import (
	"fmt"
	"net/http"
	"strings"
)

// Initialize the routes for a homepage of every location
// the route to the location is its ID
func (s *Server) initLocationHomepageRouter() {

	// Pull all locations from the database
	names, err := s.repo.GetLocationNames()
	if err != nil {
		panic(err)
	}

	// Iterate over location names and create
	for _, name := range names {
		s.router.Route(Path(processLocationsNameToRoute(name)))["GET"] = http.HandlerFunc(s.handleLocationHomepage)
	}
}

// Format locations name so that it can be used as a link
func processLocationsNameToRoute(name string) string {
	// Handle all umlauts
	name = strings.ReplaceAll(name, "ä", "+")
	name = strings.ReplaceAll(name, "ü", "*")
	name = strings.ReplaceAll(name, "ö", "$")
	return "/" + strings.ReplaceAll(name, " ", "_")
}

// handle the GET request from the location homepage
func (s *Server) handleLocationHomepage(w http.ResponseWriter, r *http.Request) {
	/*
		Determine which location
	*/
	fmt.Println(r.URL)
	// Get the name from the url
	locationName := r.URL.String()

	// Add whitespace
	locationName = strings.ReplaceAll(locationName, "_", " ")

	// Delete the '/'
	locationName = locationName[1:]

	// Derive ID from the
	fmt.Println(locationName, "was")
}
