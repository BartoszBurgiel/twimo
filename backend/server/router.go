package server

import (
	"net/http"
)

// Router type handles all routing
type Router struct {
	routes map[Path]Route
}

// Path represents the link
type Path string

// Method -> POST/GET
type Method string

// Route is a map of methods and according http handlers
type Route map[Method]http.Handler

// NewRouter constructor
func NewRouter() *Router {
	return &Router{
		routes: map[Path]Route{},
	}
}

// Route get route from path
func (r *Router) Route(p Path) Route {
	if r.routes[p] == nil {
		r.routes[p] = Route{}
	}
	return r.routes[p]
}
