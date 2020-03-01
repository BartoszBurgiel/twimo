package server

import (
	"fmt"
	"net/http"
	"twimo/backend/security"

	"github.com/gorilla/websocket"
)

// Websocket handler for login
func (s *Server) handleLoginWS(w http.ResponseWriter, r *http.Request) {

	// Open websocket connection
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	defer conn.Close()

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
	go echoLogin(conn, userData, stopper)
	<-stopper

	// Get user from the db
	userName := <-userData
	userPassword := <-userData

	user, err := s.repo.GetUserFromName(userName)
	if err != nil {
		fmt.Println(err)
	}

	// double check if user is not found
	if user.Name == "" {

		// Send message that states that login failed
		// type 1 -> text message
		err := conn.WriteMessage(1, []byte("0"))
		if err != nil {
			fmt.Println(err)
		}
		conn.Close()
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
		// Send message that states that login failed
		// type 1 -> text message
		err := conn.WriteMessage(1, []byte("1"))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successful login")
		fmt.Println(user)

	} else {
		// Send message that states that login failed
		// type 1 -> text message
		err := conn.WriteMessage(1, []byte("0"))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("invalid username or email")
		conn.Close()
	}

}

// Websocket echoer
func echoLogin(conn *websocket.Conn, userData chan string, stopper chan bool) {

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
