package server

import (
	"net/http"
)

func (s *Server) handleHomeScreenGET(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../server/assets/index.html")
}
