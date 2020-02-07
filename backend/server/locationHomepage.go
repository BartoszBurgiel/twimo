package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"twimo/backend/core"

	"github.com/gorilla/websocket"
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
		s.router.Route(Path(processLocationsNameToRoute(name) + "/ws"))["GET"] = http.HandlerFunc(s.handleLocationHomepageWS)
	}
}

// determine and return a location based on the url
func determineLocation(r *http.Request, s *Server) (location core.Location, err error) {
	// Get the name from the url
	locationName := r.URL.String()

	// Pull data from the database -> use processed location name
	location, err = s.repo.GetLocationFromName(processRouteToLocationName(locationName))
	if err != nil {
		fmt.Println(err)
		return location, err
	}

	// Set location's link to the url
	location.LocationLink = locationName

	// Pull all comments of the location
	comments, err := s.repo.GetCommentsOfLocation(location.ID)
	if err != nil {
		fmt.Println(err)
		return location, err
	}

	// Assign comments to locations' member
	location.Comments = comments
	return location, err
}

// handle the GET request from the location homepage
func (s *Server) handleLocationHomepage(w http.ResponseWriter, r *http.Request) {
	/*
		Determine which location
	*/
	locationData, err := determineLocation(r, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Host file
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

// websocket handler for location homepage
func (s *Server) handleLocationHomepageWS(w http.ResponseWriter, r *http.Request) {

	locationData, err := determineLocation(r, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Execute only when not on /ws
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	// Define the ubgrader that handles read and write buffer
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Make connection
	conn, err := upgrader.Upgrade(w, r, nil)

	// Error handling
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Start goroutine echoing the websocket
	go echoLocationHomepage(conn, s, locationData)
}

// Websocket echoer
func echoLocationHomepage(conn *websocket.Conn, s *Server, locationData core.Location) {
	// Schedule closure of the connection
	defer conn.Close()

	// Write the user to the ws
	if err := conn.WriteJSON(locationData); err != nil {
		fmt.Println(err)
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
	locationName = strings.ReplaceAll(locationName, "_", " ")

	// Split after slashes
	locationNameSplitted := strings.Split(locationName, "/")

	// If got from websocket -> if last element == "ws"
	if locationNameSplitted[len(locationNameSplitted)-1] == "ws" {
		return locationNameSplitted[len(locationNameSplitted)-2]
	}

	return locationNameSplitted[len(locationNameSplitted)-1]
}
