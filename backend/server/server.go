package server

import (
	"fmt"
	"net/http"
	"strings"
	"twimo/backend/repository"
)

// Server struct
type Server struct {
	router *Router
	repo   repository.Repository
}

// NewServer returns new server
func NewServer(r repository.Repository) (*Server, error) {

	s := &Server{
		repo: r,
	}

	s.router = NewRouter()

	return s, s.init()
}

// Innitialize the server and add routing
func (s *Server) init() error {

	/*
	  MAIN ROUTES
	*/

	// Route for the home screen
	s.router.Route("/")["GET"] = http.HandlerFunc(s.handleHomeScreenGET)

	// Route for the login page
	s.router.Route("/login")["GET"] = http.HandlerFunc(s.handleLoginGET)

	// Route for sign in page
	s.router.Route("/new")["GET"] = http.HandlerFunc(s.handleSigninGET)

	// Route for the list of all locations
	s.router.Route("/listscreen")["GET"] = http.HandlerFunc(s.handleListScreenGET)

	// Init location home pages
	s.initLocationHomepageRouter()

	/* TODO: user routing??? security???
	- how to determine if user isn't just following the link?
	- hash?
		- ssh512?
		- kihmo
			- user to hash function already prepared
		- louis stores user data in local storage + user key
	*/
	s.initUserPageRouter()

	return nil
}

// ServeHTTP to the server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if !prefixChecker(url, "/style", "/favicon.ico") {
		p := Path(r.URL.Path)
		m := Method(r.Method)

		fmt.Println(p, m)

		s.router.Route(p)[m].ServeHTTP(w, r)
		return
	}

	s.handleGETAssets(w, r)
}

// If assets are called -> host assets
func (s Server) handleGETAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("../server/assets/")).ServeHTTP(w, r)
}

// prefixChecker checks if any of given prefixes is in the url
func prefixChecker(url string, prefix ...string) bool {
	out := false
	for _, p := range prefix {
		if strings.HasPrefix(url, p) {
			out = true
		}
	}
	return out
}
