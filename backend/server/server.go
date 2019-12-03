package server

import (
	"fmt"
	"kanbanBoard/engine"
	"net/http"
	"twimo/backend/repository"

	"github.com/gorilla/websocket"
)

// Server struct
type Server struct {
	router *Router
	engine *engine.Engine
	repo   repository.Repository
}

// NewServer returns new server
func NewServer(r repository.Repository) (*Server, error) {
	newEngine := &engine.Engine{}

	s := &Server{
		engine: newEngine,
		repo:   r,
	}

	s.router = NewRouter()

	return s, s.init()
}

// Innitialize the server and add routing
func (s *Server) init() error {

	// Add route for main page
	s.router.Route("/")["GET"] = http.HandlerFunc(s.handleIndexGET)

	return nil
}

// Create a buffer that will store data to send
// via websocket
var upgrader = websocket.Upgrader{
	// Bits
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Serve HTTP protocol
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get Path
	p := Path(r.URL.Path)

	// Get Method
	m := Method(r.Method)

	// Log requests
	fmt.Println("Path", p)
	fmt.Println("Method", m)

	// Call function
	s.router.Route(p)[m].ServeHTTP(w, r)
}
