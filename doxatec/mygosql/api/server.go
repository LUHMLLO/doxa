package api

import (
	"MyGoSql/storage"
	"net/http"
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
	http.HandleFunc("/", s.Router)
	return http.ListenAndServe(s.listenAddress, nil)
}
