package server

import (
	"fmt"
	"net/http"
	"twimo/backend/security"
)

// Handle and serve admin page
func (s *Server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	// Host login file
	http.ServeFile(w, r, "../security/assets/login.html")
}

func (s *Server) handleAdminUsersGET(w http.ResponseWriter, r *http.Request) {
	ok := security.AuthAdminSession(w, r)

	fmt.Println(ok)
}

func (s *Server) handleAdminLocations(w http.ResponseWriter, r *http.Request) {
	ok := security.AuthAdminSession(w, r)

	fmt.Println(ok)
}
