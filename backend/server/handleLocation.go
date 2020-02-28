package server

import (
	"fmt"
	"net/http"
	"strings"
	"twimo/backend/core"

	"github.com/gorilla/websocket"
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

// determine and return a location based on the url
func determineLocation(r *http.Request, s *Server) (location core.Location, err error) {
	// Get the name from the url
	locationName := r.URL.String()

	// Pull data from the database -> use processed location name
	location, err = s.repo.GetLocation(processRouteToLocationID(locationName))
	if err != nil {
		fmt.Println(err)
		return location, err
	}

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

// websocket handler for location homepage
func (s *Server) handleLocationHomepage(w http.ResponseWriter, r *http.Request) {

	locationData, err := determineLocation(r, s)
	if err != nil {
		fmt.Println(err)
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

// Format routes and return the valid location name
func processRouteToLocationID(locationRoute string) string {
	// Split after slashes
	locationNameSplitted := strings.Split(locationRoute, "/")

	// return last element [id]
	return locationNameSplitted[len(locationNameSplitted)-1]
}
