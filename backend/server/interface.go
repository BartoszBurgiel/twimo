package server

import "net/http"

/*
	every screen has two methods
	-GET request handler
		- open WS connection and send initial jsons
	-WS handler
		- main communication

*/

// server interface defines which requests the server
// needs to be able to handle
type server interface {
	// (if ever needed)
	handleHomeScreenGET(w http.ResponseWriter, r *http.Request)

	// handle the GET request of the login page -> establish ws connection
	handleLoginGET(w http.ResponseWriter, r *http.Request)
	// websocket communication on the login page
	handleLoginWS(w http.ResponseWriter, r *http.Request)

	// handle the get request of the signin page -> establish ws connection
	handleSigninGET(w http.ResponseWriter, r *http.Request)
	// websocket communication on the signin page
	handleSigninWS(w http.ResponseWriter, r *http.Request)

	// handle the GET request of the listscreen page -> establish ws connection
	handleListScreenGET(w http.ResponseWriter, r *http.Request)
	// websocket communication on the listscreen page
	handleListScreenWS(w http.ResponseWriter, r *http.Request)

	// handle the GET request of the location page -> establish ws connection
	handleLocationGET(w http.ResponseWriter, r *http.Request)
	// websocket communication on the location page
	handleLocationWS(w http.ResponseWriter, r *http.Request)

	// handle the GET request of the user page -> establish ws connection
	handleUserGET(w http.ResponseWriter, r *http.Request)
	// websocket communication on the user page
	handleUserWS(w http.ResponseWriter, r *http.Request)

	// handle the GET request of the new location page -> establish ws connection
	handleNewGET(w http.ResponseWriter, r *http.Request)
	// websocket communication on the new location page
	handleNewWS(w http.ResponseWriter, r *http.Request)
}
