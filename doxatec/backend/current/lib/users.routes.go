package lib

import (
	"encoding/json"
	"fmt"
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

	users, err := s.store.users_readTable()

	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading table: %v", err))
		return
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
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading request body: %v", err))
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

	_, err = s.store.users_beforeInsert(user)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error before insertion: %v", err))
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error hashing password: %v", err))
		return
	}

	err = s.store.users_insert(user)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error during insertion: %v", err))
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
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading route params: %v", err))
		return
	}

	user, err := s.store.users_read(id)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading user: %v", err))
		return
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
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading route params: %v", err))
		return
	}

	createUserReq := &CreateUserRequest{}
	err = json.NewDecoder(r.Body).Decode(&createUserReq)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading request body: %v", err))
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

	if user.Username != "" {
		if err = s.store.users_update(id, "username", user.Username); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating username: %v", err))
			return
		}
	}

	if user.Password != "" {
		user.Password, err = GeneratehashPassword(user.Password)
		if err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error hashing password: %v", err))
			return
		}

		if err = s.store.users_update(id, "password", user.Password); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating password: %v", err))
			return
		}
	}

	if user.Avatar != "" {
		if err = s.store.users_update(id, "avatar", user.Avatar); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating avatar: %v", err))
			return
		}
	}

	if user.Name != "" {
		if err = s.store.users_update(id, "name", user.Name); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating name: %v", err))
			return
		}
	}

	if user.Email != "" {
		if err = s.store.users_update(id, "email", user.Email); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating email: %v", err))
			return
		}
	}

	if user.Phone != "" {
		if err = s.store.users_update(id, "phone", user.Phone); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating phone: %v", err))
			return
		}
	}

	if user.Role != "" {
		if err = s.store.users_update(id, "role", user.Role); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating role: %v", err))
			return
		}
	}

	if err = s.store.users_update(id, "modified", time.Now().UTC()); err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error updating modified: %v", err))
		return
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
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading route params: %v", err))
		return
	}

	err = s.store.users_delete(id)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error deleting user: %v", err))
		return
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"deleted": id})
}
