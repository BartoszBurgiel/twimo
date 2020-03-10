package server

import (
	"fmt"
	"net/http"
	"text/template"
	"twimo/backend/core"
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
	security.CheckKillSession(w, r)

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

	// Temporary object that will be parsed into the template
	type tempUserWithComments struct {
		Name     string
		Email    string
		ID       string
		Comments []core.Comment
	}

	// Build temp users with comments
	tempUsers := []tempUserWithComments{}
	for _, user := range users {

		// Get user's comments
		comments, err := s.repo.GetUsersComments(user.ID)
		if err != nil {
			fmt.Println(err)
			return
		}

		tempUsers = append(tempUsers, tempUserWithComments{
			Name:     user.Name,
			Email:    user.Email,
			ID:       user.ID,
			Comments: comments,
		})

	}

	// Path reassigned to variable p to shorten the code
	p := pathToAdminAssets

	temp := template.Must(template.ParseFiles(p+"users.html", p+"user.html", p+"comment.html"))
	err = temp.Execute(w, tempUsers)
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
	// Check if the admin has logged in
	if !security.AuthAdminSession(w, r) {
		w.Write([]byte(`You need to log in!`))
		return
	}

	// Fetch post variables
	deleteComment := r.FormValue("deleteComment")
	deleteLocation := r.FormValue("deleteLocation")

	// determine what to do

	// id delete comment was clicked
	// If delete comment has been clicked
	if deleteComment != "" {
		fmt.Println("commentID to delete: ", deleteComment)
		err := s.repo.DeleteComment(deleteComment)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("comment has been deleted")
	} else if deleteLocation != "" {
		fmt.Println("locationID to delete: ", deleteComment)
		err := s.repo.DeleteLocation(deleteComment)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("location has been deleted")
	}

	// reopen the page
	s.handleAdminLocations(w, r)
}

func (s *Server) handleAdminUsersPOST(w http.ResponseWriter, r *http.Request) {
	security.CheckKillSession(w, r)
	// Check if the admin has logged in
	if !security.AuthAdminSession(w, r) {
		w.Write([]byte(`You need to log in!`))
		return
	}

	// Check post values
	deleteComment := r.FormValue("deleteComment")
	deleteUser := r.FormValue("deleteUser")

	// If delete comment has been clicked
	if deleteComment != "" {
		fmt.Println("commentID to delete: ", deleteComment)
		err := s.repo.DeleteComment(deleteComment)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("comment has been deleted")
		// if delete user has been clicked
	} else if deleteUser != "" {
		fmt.Println("userID to delete: ", deleteComment)
		err := s.repo.DeleteUser(deleteUser)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("user has been deleted")
	}
	s.handleAdminUsersGET(w, r)
}
