package server

import (
	"fmt"
	"net/http"
	"text/template"
	"twimo/backend/security"
)

var pathToAdminAssets = "../server/assets/admin/"

// Handle and serve admin page
func (s *Server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	// Host login file
	http.ServeFile(w, r, "../security/assets/login.html")
}

func (s *Server) handleAdminPOST(w http.ResponseWriter, r *http.Request) {
	security.AdminLogin(w, r)

	// Check if the admin has logged in
	if !security.AuthAdminSession(w, r) {
		w.Write([]byte(`
		Sorry, something went wrong or you entered the wrong password ¯\_(ツ)_/¯
		`))
		return
	}

	http.ServeFile(w, r, "../server/assets/admin/index.html")
}

func (s *Server) handleAdminUsersGET(w http.ResponseWriter, r *http.Request) {
	// Check if the admin has logged in
	if !security.AuthAdminSession(w, r) {
		w.Write([]byte(`You need to log in!`))
		return
	}

	// Get all users
	users, err := s.repo.GetAllUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Path reassigned to variable p to shorten the code
	p := pathToAdminAssets

	temp := template.Must(template.ParseFiles(p+"users.html", p+"user.html"))
	err = temp.Execute(w, users)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *Server) handleAdminLocations(w http.ResponseWriter, r *http.Request) {
	// Check if the admin has logged in
	if !security.AuthAdminSession(w, r) {
		w.Write([]byte(`You need to log in!`))
		return
	}

	// get all locations
	locations, err := s.repo.GetAllLocations()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Path reassigned to variable p to shorten the code
	p := pathToAdminAssets

	temp := template.Must(template.ParseFiles(p+"locations.html", p+"location.html", p+"comment.html"))
	err = temp.Execute(w, locations)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *Server) handleAdminLocationsPOST(w http.ResponseWriter, r *http.Request) {
	security.CheckKillSession(w, r)
}

func (s *Server) handleAdminUsersPOST(w http.ResponseWriter, r *http.Request) {
	security.CheckKillSession(w, r)
}
