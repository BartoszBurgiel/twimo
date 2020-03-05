package server

import "net/http"

// Handle and serve admin page
func (s *Server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	// Host index.html
	http.ServeFile(w, r, "../server/assets/admin/index.html")
}

func (s *Server) handleAdminUsers(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) handleAdminLocations(w http.ResponseWriter, r *http.Request) {

}
