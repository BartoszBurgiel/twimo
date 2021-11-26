package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// handle the GET request for the location list page
func (s *Server) handleListScreenGET(w http.ResponseWriter, r *http.Request) {

	// Open websocket connection
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	defer conn.Close()
	// GET MESSAGE AND DETERMINE HOW TO SORT
	sorting := make(chan string)

	// Listen to the websocket and fetch the sorting clue
	go listenForSortingCriteria(conn, sorting)

	for {
		fmt.Println("going")

		// Get all locations from the database
		allLocations, err := s.repo.GetLocationsForList(<-sorting)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Send locations as an array of locations
		if err := conn.WriteJSON(allLocations); err != nil {
			fmt.Println(err)
		}

	}
}

// Listen to the websocket and establish the sorting criteria
func listenForSortingCriteria(conn *websocket.Conn, sorting chan string) {
	for {
		// Get message
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message", err)
			return
		}

		// Convert msg to string
		sorting <- string(msg)
	}
}
