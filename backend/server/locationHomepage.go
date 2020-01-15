package server

import (
	"fmt"
	"html/template"
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

// handle the GET request from the location homepage
func (s *Server) handleLocationHomepage(w http.ResponseWriter, r *http.Request) {
	/*
		Determine which location
	*/

	// Get the name from the url
	locationName := r.URL.String()

	// Add whitespace
	locationName = processRouteToLocationName(locationName)

	// Pull data from the database
	locationData, err := s.repo.GetLocationFromName(locationName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Pull all comments of the location
	comments, err := s.repo.GetCommentsOfLocation(locationData.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Assign comments to locations' member
	locationData.Comments.Comments = comments

	// Send the data to the client via file
	temp, err := template.New("location").ParseFiles("../server/assets/location.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = temp.Execute(w, locationData)
	if err != nil {
		fmt.Println(err)
		return
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

// Format routes and return the valid location name
func processRouteToLocationName(locationName string) string {
	// Handle all umlauts
	locationName = strings.ReplaceAll(locationName, "+", "ä")
	locationName = strings.ReplaceAll(locationName, "*", "ü")
	locationName = strings.ReplaceAll(locationName, "$", "ö")

	// Remove slash(es)
	locationName = strings.ReplaceAll(locationName, "/", "")

	return strings.ReplaceAll(locationName, "_", " ")
}
