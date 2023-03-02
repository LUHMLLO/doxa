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

	router.HandleFunc("/auth/signup", s.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/signin", s.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/check-admin", IsAuthorized(s.CheckAdmin)).Methods("GET", "OPTIONS")
	router.HandleFunc("/auth/check-role", IsAuthorized(s.CheckRole)).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/users", s.Handle_allUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/users/create", s.Handle_insertUsers).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/users/read/{id}", s.Handle_readUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/users/update/{id}", s.Handle_updateUsers).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/users/delete/{id}", s.Handle_deleteUsers).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/devices", s.Handle_allDevices).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/devices/create", s.Handle_insertDevices).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/devices/read/{id}", s.Handle_readDevices).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/devices/update/{id}", s.Handle_updateDevices).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/devices/delete/{id}", s.Handle_deleteDevices).Methods("DELETE", "OPTIONS")

	log.Println("Doxatec server running on port:", s.listenAddress)
	log.Fatal(http.ListenAndServe(s.listenAddress, router))
}
