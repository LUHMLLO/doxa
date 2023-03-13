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
	router.HandleFunc("/api/auth/signature", Authenticate(s.SignedUser)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth/signout", Authenticate(s.SignOut)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/mydevices", Authenticate(s.UserDevices)).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/auth/check-admin", s.UserAdmin).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/auth/check-role", s.UserRole).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/auth/check-username", s.UserName).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/master/users/all", Authenticate(s.Route_all_users)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/master/users/insert", Authenticate(s.Route_insert_user)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/master/users/read/{id}", Authenticate(s.Route_read_user)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/master/users/update/{id}", Authenticate(s.Route_update_user)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/master/users/delete/{id}", Authenticate(s.Route_delete_user)).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/master/devices/all", Authenticate(s.Route_all_devices)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/master/devices/insert", Authenticate(s.Route_insert_device)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/master/devices/read/{id}", Authenticate(s.Route_read_device)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/master/devices/update/{id}", Authenticate(s.Route_update_device)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/master/devices/delete/{id}", Authenticate(s.Route_delete_device)).Methods("DELETE", "OPTIONS")

	log.Println("Doxatec server running on port:", s.listenAddress)
	log.Fatal(http.ListenAndServe(s.listenAddress, router))
}
