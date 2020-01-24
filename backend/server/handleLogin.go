package server

import (
	"fmt"
	"net/http"
	"twimo/backend/security"

	"github.com/gorilla/websocket"
)

// send login.html file to the client
func (s *Server) handleLoginGET(w http.ResponseWriter, r *http.Request) {
	// Serve the file
	http.ServeFile(w, r, "../server/assets/login.html")
}

// Websocket handler for login
func (s *Server) handleLoginWS(w http.ResponseWriter, r *http.Request) {

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

	// Channel for user email and password
	userData := make(chan string, 2)

	// stopper is used to stop and start the program if the user data are put in
	stopper := make(chan bool)

	// Start goroutine echoing the websocket
	go echoLogin(conn, s, userData, stopper)
	<-stopper

	// Get user from the db
	fmt.Println("here we go")
	userName := <-userData
	userPassword := <-userData

	user, err := s.repo.GetUserFromName(userName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// double check if user is not found
	if user.Name == "" {
		fmt.Println("no user found")
		return
	}

	// check password
	ok, err := security.Password(user, userPassword, s.repo)
	if err != nil {
		fmt.Println(err)
		return
	}

	if ok {
		fmt.Println("Successful login")
		fmt.Println(user)
	} else {
		fmt.Println("invalid username or email")
	}

}

// Websocket echoer
func echoLogin(conn *websocket.Conn, s *Server, userData chan string, stopper chan bool) {
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
		userData <- string(msg)

		// Log message
		fmt.Println("Got message: ", string(msg))

		if len(userData) == 2 {
			stopper <- true
			return
		}
	}
}
