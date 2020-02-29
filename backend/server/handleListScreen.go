package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// handle the GET request for the location list page
func (s *Server) handleListScreenGET(w http.ResponseWriter, r *http.Request) {

	// GET MESSAGE AND DETERMINE HOW TO SORT

	// Get all locations from the database
	allLocations, err := s.repo.GetLocationsForList("rating")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Open websocket connection
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	defer conn.Close()

	// Iterate over locations and send JSONS
	for _, location := range allLocations {
		if err := conn.WriteJSON(location); err != nil {
			fmt.Println(err)
			return
		}
	}

	// Close the connection
	conn.Close()
}
