package server

import (
	"net/http"
)

// Handle GET request on /
func (s *Server) handleIndexGET(w http.ResponseWriter, r *http.Request) {

	// Serve index.html file
	http.ServeFile(w, r, "../server/assets/index.html")

	s.repo.GetUser("hahah")
	s.repo.GetUser("69696")
}
