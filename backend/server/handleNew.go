package server

import (
	"fmt"
	"net/http"
	"strconv"
	"twimo/backend/core"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Handle the websocket connection and addition of a new location
func (s *Server) handleNewWS(w http.ResponseWriter, r *http.Request) {

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

	// Multithreading variables
	locationChannel := make(chan tempLocationForWS, 1)
	stopper := make(chan bool)

	// Listen for location information
	go listenForLocationInformation(conn, stopper, locationChannel)
	<-stopper

	// Fetch location
	tempLocation := <-locationChannel

	// Convert location
	location := tempLocation.toLocation()

	// Add location to database
	err = s.repo.AddLocation(location)
	if err != nil {
		fmt.Println(err)
		// Send message that user has been added to the database
		// type 1 -> text message
		err = conn.WriteMessage(1, []byte("0"))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Send confirmation
	// Send message that user has been added to the database
	// type 1 -> text message
	err = conn.WriteMessage(1, []byte("1"))
	if err != nil {
		fmt.Println(err)
	}

	conn.Close()
}

// listen for the location information that will be sent
// via the websocket
func listenForLocationInformation(conn *websocket.Conn, stopper chan bool, location chan tempLocationForWS) {

	// temporary location
	temp := tempLocationForWS{}
	for {
		err := conn.ReadJSON(&temp)
		if err != nil {
			fmt.Println(err)
			// Send message that user has been added to the database
			// type 1 -> text message
			err = conn.WriteMessage(1, []byte("0"))
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		if temp.Name != "" {
			location <- temp
			stopper <- true
			return
		}
	}
}

// temporary struct with different fields to
// easen the websocket communication
type tempLocationForWS struct {
	Name string

	// Store geological x and y coordinates
	Coords struct {
		X string
		Y string
	}

	Desc string

	Price string

	Features core.Features
}

// Convert the temporary struct to location
// and assign ID
func (t tempLocationForWS) toLocation() (loc core.Location) {
	loc.Name = t.Name
	loc.Desc = t.Desc

	// Convert strings to floats
	x, _ := strconv.ParseFloat(t.Coords.X, 64)
	y, _ := strconv.ParseFloat(t.Coords.Y, 64)

	loc.Features = t.Features
	loc.Coords = struct {
		X float64
		Y float64
	}{
		X: x,
		Y: y,
	}

	loc.Price, _ = strconv.Atoi(t.Price)

	// Assign ID
	loc.ID = uuid.New().String()

	return loc
}
