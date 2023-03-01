package lib

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddress string
	store         Storage
}

func NewServer(listenAddress string, store Storage) *Server {
	return &Server{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *Server) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", s.Users_Handler).Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
	router.HandleFunc("/api/devices", s.Devices_Handler).Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")

	log.Println("Doxatec server running on port:", s.listenAddress)
	log.Fatal(http.ListenAndServe(s.listenAddress, router))
}
