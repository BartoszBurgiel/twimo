package server

import (
	"fmt"
	"net/http"
	"twimo/backend/security"

	"github.com/gorilla/websocket"
)

func (s *Server) handleUserWS(w http.ResponseWriter, r *http.Request) {

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

	// Stores user data needed to auth
	tempUserData := make(chan security.UserData, 1)
	stopper := make(chan bool)

	// Wait and listen for the messages
	go listenForUserPageData(conn, tempUserData, stopper)
	<-stopper
	fmt.Println("ok")

	// Fetch tempUserData
	userDataToCheck := <-tempUserData

	// Check
	ok := security.IsHashAndUsernameValid(userDataToCheck, s.repo)
	if !ok {
		// Send message that user needs to relog because key is
		// not valid
		// type 1 -> text message
		err = conn.WriteMessage(1, []byte("0"))
		if err != nil {
			fmt.Println(err)
		}

		conn.Close()
		return
	}

	// Send message that user has successfully logged in
	// type 1 -> text message
	err = conn.WriteMessage(1, []byte("1"))
	if err != nil {
		fmt.Println(err)
	}

	/* Send user's data */
	user, err := s.repo.GetUserFromName(userDataToCheck.Username)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Destroy password
	user.Password = ""

	// Send to the client
	// Write the user to the ws
	if err := conn.WriteJSON(user); err != nil {
		fmt.Println(err)
	}

	// Send location
	location, err := s.repo.GetLocationCard(user.FavLocation)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send to the client
	// Write the user to the ws
	if err := conn.WriteJSON(location); err != nil {
		fmt.Println(err)
	}
}

// listen to the connection and gather the data
func listenForUserPageData(conn *websocket.Conn, tempUserData chan security.UserData, stopper chan bool) {

	for {

		temp := security.UserData{
			Username: "",
			Key:      "",
		}
		// fetch JSON from fronend
		err := conn.ReadJSON(&temp)
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}

		// If temp has been filled
		if temp.Username != "" && temp.Key != "" {
			// Write to channel
			tempUserData <- temp

			// Relieve stopper
			stopper <- true

			// Escape from function
			return
		}

	}
}
