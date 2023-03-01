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

	router.HandleFunc("/api/users", s.ReadAllUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/create/user", s.InsertUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/read/user/{id}", s.ReadUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/update/user/{id}", s.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete/user/{id}", s.DeleteUser).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/devices", s.ReadAllDevices).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/create/device", s.InsertDevice).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/read/device/{id}", s.ReadDevice).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/update/device/{id}", s.UpdateDevice).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete/device/{id}", s.DeleteDevice).Methods("DELETE", "OPTIONS")

	log.Println("Doxatec server running on port:", s.listenAddress)
	log.Fatal(http.ListenAndServe(s.listenAddress, router))
}
