package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	user := s.store.GetUser(uuid.New())
	json.NewEncoder(w).Encode(user)
}

func (s *Server) HandleGetOwners(w http.ResponseWriter, r *http.Request) {
	owner := s.store.GetOwner(uuid.New())
	json.NewEncoder(w).Encode(owner)
}
func (s *Server) HandleGetDevices(w http.ResponseWriter, r *http.Request) {
	device := s.store.GetDevice(uuid.New())
	json.NewEncoder(w).Encode(device)
}
