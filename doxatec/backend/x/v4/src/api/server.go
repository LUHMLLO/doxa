package api

import (
	"doxatec/types"
	"doxatec/utils"
	"log"
	"net/http"
	"os"

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
	os.Setenv("JWT_SECRET", "D@x@dm1n_JWT_Secret")

	router := mux.NewRouter()

	router.HandleFunc("/admin/users", utils.MakeHttpHandleFunc(s.Handle_User))
	router.HandleFunc("/api/user/{id}", utils.ProtectWithJWT(utils.MakeHttpHandleFunc(s.Handle_ReadUserById), s.store))

	router.HandleFunc("/admin/profiles", utils.MakeHttpHandleFunc(s.Handle_Profile))
	router.HandleFunc("/api/profile/{id}", utils.MakeHttpHandleFunc(s.Handle_ReadProfileById))

	router.HandleFunc("/admin/devices", utils.MakeHttpHandleFunc(s.Handle_Device))
	router.HandleFunc("/api/device/{id}", utils.MakeHttpHandleFunc(s.Handle_ReadDeviceById))

	log.Println("Doxatec server running on port:", s.listenAddress)
	log.Println("Available api endpoints:")
	log.Printf("http://localhost%s\n", s.listenAddress)
	log.Printf("http://localhost%s/admin/users\n", s.listenAddress)
	log.Printf("http://localhost%s/api/user/{id}\n", s.listenAddress)
	log.Printf("http://localhost%s/admin/profiles\n", s.listenAddress)
	log.Printf("http://localhost%s/api/profile/{id}\n", s.listenAddress)
	log.Printf("http://localhost%s/admin/devices\n", s.listenAddress)
	log.Printf("http://localhost%s/api/device/{id}\n", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)
}
