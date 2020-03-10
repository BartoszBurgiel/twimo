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
	s.router.Route("/login")["GET"] = http.HandlerFunc(s.handleLoginWS)

	// Route for sign in page
	s.router.Route("/signup")["GET"] = http.HandlerFunc(s.handleSignupGET)

	// Route for the list of all locations
	s.router.Route("/listscreen")["GET"] = http.HandlerFunc(s.handleListScreenGET)

	// Route for the userpage
	s.router.Route("/user")["GET"] = http.HandlerFunc(s.handleUserWS)

	// Route for the new location
	s.router.Route("/new")["GET"] = http.HandlerFunc(s.handleNewWS)

	/*
		ADMIN SECTION
	*/
	s.router.Route("/admin")["GET"] = http.HandlerFunc(s.handleAdmin)
	s.router.Route("/admin")["POST"] = http.HandlerFunc(s.handleAdminPOST)

	s.router.Route("/admin/locations")["GET"] = http.HandlerFunc(s.handleAdminLocations)
	s.router.Route("/admin/locations")["POST"] = http.HandlerFunc(s.handleAdminLocationsPOST)
	s.router.Route("/admin/users")["GET"] = http.HandlerFunc(s.handleAdminUsersGET)
	s.router.Route("/admin/users")["POST"] = http.HandlerFunc(s.handleAdminUsersPOST)

	// Init location home pages
	s.initLocationHomepageRouter()

	return nil
}

// ServeHTTP to the server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if !prefixChecker(url, "/admin/style", "/favicon.ico") {
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
	http.FileServer(http.Dir("../server/assets")).ServeHTTP(w, r)
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
