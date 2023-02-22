package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddress string
	store         Storage
}

func NewApiServer(listenAddress string, store Storage) *ApiServer {
	return &ApiServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *ApiServer) Start() {
	router := mux.NewRouter()
	router.HandleFunc("/users", MakeHttpHandleFunc(s.handleUser))
	router.HandleFunc("/users/{id}", MakeHttpHandleFunc(s.handleUser))

	log.Println("DOXA api server ruuning on port:", s.listenAddress)
	log.Printf("http://localhost%s\n", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

func (s *ApiServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateUser(w, r)
	case "GET":
		return s.handleReadUser(w, r)
	case "PUT":
		return s.handleUpdateUser(w, r)
	case "DELETE":
		return s.handleDeleteUser(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *ApiServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleReadUser(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)

	return WriteJSON(w, http.StatusOK, &User{})
}

func (s *ApiServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
