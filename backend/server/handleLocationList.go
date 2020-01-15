package server

import (
	"fmt"
	"html/template"
	"net/http"
)

// handle the GET request for the location list page
func (s *Server) handleLocationListGET(w http.ResponseWriter, r *http.Request) {

	// Get all locations from the database
	allLocations, err := s.repo.GetLocationsAsLink()
	if err != nil {
		fmt.Println(err)
		return
	}

	temp, err := template.New("locationList").ParseFiles("../server/assets/locationList.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = temp.Execute(w, allLocations)
	if err != nil {
		fmt.Println(err)
		return
	}
}
