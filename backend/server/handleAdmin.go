package server

import (
	"fmt"
	"net/http"
	"twimo/backend/security"
)

// Handle and serve admin page
func (s *Server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	// Host index.html
	http.ServeFile(w, r, "../server/assets/admin/index.html")
}

func (s *Server) handleAdminUsersGET(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../server/assets/admin/login.html")
}

func (s *Server) handleAdminUsersPOST(w http.ResponseWriter, r *http.Request) {
	enteredPassword := r.FormValue("pword")

	fmt.Println(enteredPassword)

	// Check password
	if security.AuthAdminPassword(enteredPassword) {

		// Host all users
		serveAllUsersTemplate()

	} else {
		w.Write([]byte(`
		Das von Ihn eingegebene Passwort war leider falsch. <br>
		<a href="admin/users">Bitte versuchen Sie erneut. </a>`))
	}

}

func (s *Server) handleAdminLocations(w http.ResponseWriter, r *http.Request) {

}
