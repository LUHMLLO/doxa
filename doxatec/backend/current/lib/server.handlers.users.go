package lib

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) Handle_allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	users, err := s.store.Query_allUsers()

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users)
}

func (s *Server) Handle_insertUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	createUserReq := &CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&createUserReq)
	if err != nil {
		log.Fatalln("error in reading body: ", err.Error())
		return
	}

	user := NewUser(
		createUserReq.JWT,
		createUserReq.Username,
		createUserReq.Password,
		createUserReq.Avatar,
		createUserReq.Name,
		createUserReq.Email,
		createUserReq.Phone,
		createUserReq.Role,
	)

	_, err = s.store.Query_beforeInsertUsers(user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error hashing password: ", err.Error())
		return
	}

	err = s.store.Query_insertUsers(user)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) Handle_readUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	user, err := s.store.Query_readUsers(id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) Handle_updateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	createUserReq := &CreateUserRequest{}
	err = json.NewDecoder(r.Body).Decode(&createUserReq)
	if err != nil {
		log.Fatal(err)
	}

	user := NewUser(
		createUserReq.JWT,
		createUserReq.Username,
		createUserReq.Password,
		createUserReq.Avatar,
		createUserReq.Name,
		createUserReq.Email,
		createUserReq.Phone,
		createUserReq.Role,
	)

	err = s.store.Query_updateUsers(id, user)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"updated": id})
}

func (s *Server) Handle_deleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.store.Query_deleteUsers(id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"deleted": id})
}
