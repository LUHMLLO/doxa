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

	router.HandleFunc("/api/auth/signup", s.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/signin", s.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/signature", s.SignedUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth/signout", s.SignOut).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/mydevices", s.UserDevices).Methods("GET", "OPTIONS")

	// router.HandleFunc("/api/auth/check-admin", IsAuthorized(s.CheckAdmin)).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/auth/check-role", IsAuthorized(s.CheckRole)).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/auth/check-username", IsAuthorized(s.CheckUsername)).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/v1/users/all", s.Route_all_users).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/users/insert", s.Route_insert_user).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/users/read/{id}", s.Route_read_user).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/users/update/{id}", s.Route_update_user).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/v1/users/delete/{id}", s.Route_delete_user).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/v1/devices/all", s.Route_all_devices).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/devices/insert", s.Route_insert_device).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/devices/read/{id}", s.Route_read_device).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/devices/update/{id}", s.Route_update_device).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/v1/devices/delete/{id}", s.Route_delete_device).Methods("DELETE", "OPTIONS")

	log.Println("Doxatec server running on port:", s.listenAddress)
	log.Fatal(http.ListenAndServe(s.listenAddress, router))
}
