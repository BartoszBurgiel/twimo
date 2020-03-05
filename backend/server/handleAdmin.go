package server

import (
	"net/http"
	"twimo/backend/security"
)

// Handle and serve admin page
func (s *Server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	// Host login file
	http.ServeFile(w, r, "../security/assets/login.html")
}

func (s *Server) handleAdminPOST(w http.ResponseWriter, r *http.Request) {
	security.AdminLogin(w, r)

	if !security.AuthAdminSession(w, r) {
		w.Write([]byte(`
		Sorry, something went wrong or you entered the wrong password ¯\_(ツ)_/¯
		`))
		return
	}

	http.ServeFile(w, r, "../server/assets/admin/index.html")
}

func (s *Server) handleAdminUsersGET(w http.ResponseWriter, r *http.Request) {
	if !security.AuthAdminSession(w, r) {
		w.Write([]byte(`You need to log in!`))
		return
	}
}

func (s *Server) handleAdminLocations(w http.ResponseWriter, r *http.Request) {
	if !security.AuthAdminSession(w, r) {
		w.Write([]byte(`You need to log in!`))
		return
	}
}
