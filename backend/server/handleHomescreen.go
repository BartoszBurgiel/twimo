package server

import (
	"net/http"
)

func (s *Server) handleHomeScreenGET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<b>TWIMO</b> Webpage is currently still under construction. :)"))
}
