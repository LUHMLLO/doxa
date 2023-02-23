package api

import (
	"doxatec/types"
	"doxatec/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

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
	createUserReq := new(types.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&createUserReq); err != nil {
		return err
	}

	user := types.NewUser(createUserReq.Avatar, createUserReq.Username, createUserReq.Password)
	if err := s.store.CreateUser(user); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, user)
}

func (s *ApiServer) handleReadUser(w http.ResponseWriter, r *http.Request) error {
	users, err := s.store.ReadUsers()
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, users)
}

func (s *ApiServer) handleReadUserById(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case "GET":
		id, err := utils.GetID(r)
		if err != nil {
			return err
		}

		users, err := s.store.ReadUserByID(id)
		if err != nil {
			return err
		}

		return utils.WriteJSON(w, http.StatusOK, users)
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
	id, err := utils.GetID(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteUser(id); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]uuid.UUID{"deleted": id})
}
