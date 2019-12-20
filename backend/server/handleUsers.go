package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Handle GET request on /users
func (s *Server) handleUsersGET(w http.ResponseWriter, r *http.Request) {

	// Serve the file
	http.ServeFile(w, r, "../server/assets/users.html")
}

// Websocket handler for user
func (s *Server) handleUserWS(w http.ResponseWriter, r *http.Request) {

	// Execute only when not on /ws
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}

	// Define the ubgrader that handles read and write buffer
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Make connection
	conn, err := upgrader.Upgrade(w, r, nil)

	// Error handling
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Start goroutine echoing the websocket
	go echoUsers(conn, s)
}

// Websocket echoer
func echoUsers(conn *websocket.Conn, s *Server) {
	// Schedule closure of the connection
	defer conn.Close()

	for {

		// Get message
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message", err)
			return
		}

		// Convert msg to string
		userID := string(msg)

		// Log message
		fmt.Println("Got message: ", userID)

		// Get users
		if user, err := s.repo.GetUser(userID); err != nil {
			// Write the user to the db
			if err := conn.WriteMessage(1, []byte(fmt.Sprintf("No user with the id %s has been found in the DB", userID))); err != nil {
				fmt.Println(err)
			}
		} else {

			// Write the user to the db
			if err := conn.WriteJSON(user); err != nil {
				fmt.Println(err)
			}
		}
	}
}
