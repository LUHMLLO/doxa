package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) Router(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		s.RouteHome(w, r)
	case "/users":
		s.HandleGetUsers(w, r)
	case "/owners":
		s.HandleGetOwners(w, r)
	case "/devices":
		s.HandleGetDevices(w, r)
	default:
		fmt.Println("Too far away.")
	}
}

func (s *Server) RouteHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello world")
}
