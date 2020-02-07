package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Handle GET request on /comments
func (s *Server) handleCommentsGET(w http.ResponseWriter, r *http.Request) {

	// Serve the file
	http.ServeFile(w, r, "../server/assets/comments.html")
}

// Websocket handler for comment
func (s *Server) handleCommentsWS(w http.ResponseWriter, r *http.Request) {

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
	go echoComments(conn, s)
}

// Websocket echoer
func echoComments(conn *websocket.Conn, s *Server) {
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
		commentID := string(msg)

		// Log message
		fmt.Println("Got message: ", commentID)

		// Get comment
		if comment, err := s.repo.GetComment(commentID); err != nil {

			// Write log
			if err := conn.WriteMessage(1, []byte(fmt.Sprintf("No user with the id %T has been found in the DB", comment))); err != nil {
				fmt.Println(err)
			}
		} else {

			// Write the user to the ws
			if err := conn.WriteJSON(comment); err != nil {
				fmt.Println(err)
			}
		}
	}
}
