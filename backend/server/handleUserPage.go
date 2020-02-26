package server

import (
	"fmt"
	"net/http"
)

func (s *Server) handleUserGET(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) handleUserWS(w http.ResponseWriter, r *http.Request) {

}

// Create a route for every user
func (s *Server) initUserPageRouter() {
	// Get all users' IDs
	IDs, err := s.repo.GetAllUsersIDs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, id := range IDs {
		s.router.Route(Path(id))["GET"] = http.HandlerFunc(s.handleUserWS)
	}
}
