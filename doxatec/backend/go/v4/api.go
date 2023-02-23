package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
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
	router.HandleFunc("/users/{id}", MakeHttpHandleFunc(s.handleReadUserById))

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
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *ApiServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserReq := new(CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&createUserReq); err != nil {
		return err
	}

	user := NewUser(createUserReq.Avatar, createUserReq.Username, createUserReq.Password)
	if err := s.store.CreateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

func (s *ApiServer) handleReadUser(w http.ResponseWriter, r *http.Request) error {
	users, err := s.store.ReadUsers()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, users)
}

func (s *ApiServer) handleReadUserById(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case "GET":
		id, err := getID(r)
		if err != nil {
			return err
		}

		users, err := s.store.ReadUserByID(id)
		if err != nil {
			return err
		}

		return WriteJSON(w, http.StatusOK, users)
	case "PUT":
		return s.handleUpdateUser(w, r)
	case "DELETE":
		return s.handleDeleteUser(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *ApiServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteUser(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]uuid.UUID{"deleted": id})
}
