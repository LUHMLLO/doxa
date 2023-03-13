package lib

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) Route_all_users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	users, err := s.store.Query_read_all_users_from_table()

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users)
}

func (s *Server) Route_insert_user(w http.ResponseWriter, r *http.Request) {
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
		createUserReq.Username,
		createUserReq.Password,
		createUserReq.Avatar,
		createUserReq.Name,
		createUserReq.Email,
		createUserReq.Phone,
		createUserReq.Role,
	)

	_, err = s.store.Query_before_insert_user(user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error hashing password: ", err.Error())
		return
	}

	err = s.store.Query_insert_user(user)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) Route_read_user(w http.ResponseWriter, r *http.Request) {
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

	user, err := s.store.Query_read_user_where_column("id", id.String())
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) Route_update_user(w http.ResponseWriter, r *http.Request) {
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
		createUserReq.Username,
		createUserReq.Password,
		createUserReq.Avatar,
		createUserReq.Name,
		createUserReq.Email,
		createUserReq.Phone,
		createUserReq.Role,
	)

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error hashing password: ", err.Error())
		return
	}

	if user.Username != "" {
		if err = s.store.Query_update_user_column_where_ID(id, "username", user.Username); err != nil {
			log.Fatal(err)
		}
	}

	if user.Password != "" {
		if err = s.store.Query_update_user_column_where_ID(id, "password", user.Password); err != nil {
			log.Fatal(err)
		}
	}

	if user.Avatar != "" {
		if err = s.store.Query_update_user_column_where_ID(id, "avatar", user.Avatar); err != nil {
			log.Fatal(err)
		}
	}

	if user.Name != "" {
		if err = s.store.Query_update_user_column_where_ID(id, "name", user.Name); err != nil {
			log.Fatal(err)
		}
	}

	if user.Email != "" {
		if err = s.store.Query_update_user_column_where_ID(id, "email", user.Email); err != nil {
			log.Fatal(err)
		}
	}

	if user.Phone != "" {
		if err = s.store.Query_update_user_column_where_ID(id, "phone", user.Phone); err != nil {
			log.Fatal(err)
		}
	}

	if user.Role != "" {
		if err = s.store.Query_update_user_column_where_ID(id, "role", user.Role); err != nil {
			log.Fatal(err)
		}
	}

	if err = s.store.Query_update_user_column_where_ID(id, "modified", time.Now().UTC()); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"updated": id})
}

func (s *Server) Route_delete_user(w http.ResponseWriter, r *http.Request) {
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

	_, err = s.store.Query_delete_user_where_ID(id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"deleted": id})
}
