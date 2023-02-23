package api

import (
	"doxatec/types"
	"doxatec/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddress string
	store         types.Storage
}

func NewServer(listenAddress string, store types.Storage) *Server {
	return &Server{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *Server) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/users", utils.MakeHttpHandleFunc(s.Handle_User))
	router.HandleFunc("/users/{id}", utils.MakeHttpHandleFunc(s.Handle_ReadUserById))

	log.Println("Doxatec server running on port:", s.listenAddress)
	log.Println("Available api endpoints:")
	log.Printf("http://localhost%s\n", s.listenAddress)
	log.Printf("http://localhost%s/users\n", s.listenAddress)
	log.Printf("http://localhost%s/users/{id}\n", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)
}
