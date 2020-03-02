package server

import (
	"fmt"
	"net/http"
	"twimo/backend/core"
	"twimo/backend/security"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// handleSigninGET recieves the data from the frontend and
// creates new user entry
func (s *Server) handleSigninGET(w http.ResponseWriter, r *http.Request) {

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
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Channel for username, email and password
	userData := make(chan string, 3)
	uData := []string{}

	// stopper is used to stop and start the program if the user data are put in
	stopper := make(chan bool)

	// Start goroutine echoing the websocket
	go listenForSigninData(conn, userData, stopper)
	<-stopper

	// retrieve information from the channel
	for i := 0; i < 3; i++ {
		uData = append(uData, <-userData)
	}
	fmt.Println(uData)
	// Check if user exists in the database
	ok, _ := s.repo.UserExists(uData[0])
	if ok {

		// Send message that user is already in the database
		// type 1 -> text message
		err := conn.WriteMessage(1, []byte("0"))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	/* Add user to database */

	// Build user struct
	tempUser := core.User{
		Name:        uData[0],
		Email:       uData[1],
		Password:    uData[2],
		FavLocation: "",
		ID:          uuid.New().String(),
	}

	// add tempUser to database
	err = s.repo.AddUser(tempUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User has been added to the database")
	// Send message that user has been added to the database
	// type 1 -> text message
	err = conn.WriteMessage(1, []byte("1"))
	if err != nil {
		fmt.Println(err)
	}

	// Send the hashed user key

	// Generate key
	key := tempUser.ToChecksum()

	// assemble json
	temp := security.UserData{
		Username: tempUser.Name,
		Key:      key,
	}

	// Send message
	err = conn.WriteJSON(temp)
	if err != nil {
		fmt.Println(err)
	}

	// close connection
	conn.Close()
}

// listenForSigninData listens to the port and waits for the user to
// enter his credentials
func listenForSigninData(conn *websocket.Conn, userData chan string, stopper chan bool) {
	for {

		// Get message
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message", err)
			return
		}

		// Convert msg to string
		userData <- string(msg)

		if len(userData) == 3 {
			stopper <- true
			return
		}
	}
}
