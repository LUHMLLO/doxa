package api

import (
	"MyGoSql/storage"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddress string
	store         storage.Storage
}

func NewServer(listenAddress string, store storage.Storage) *Server {
	return &Server{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *Server) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/users", s.HandleGetUsers)
	router.HandleFunc("/owners", s.HandleGetOwners)
	router.HandleFunc("/devices", s.HandleGetDevices)
	return http.ListenAndServe(s.listenAddress, nil)
}
