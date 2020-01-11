package server

import (
	"fmt"
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

	return "/" + strings.ReplaceAll(name, " ", "")
}

func (s *Server) handleLocationHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL, "helllo")
}
