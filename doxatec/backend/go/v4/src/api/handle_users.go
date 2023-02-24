package api

import (
	"doxatec/types"
	"doxatec/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) Handle_User(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.Handle_CreateUser(w, r)
	case "GET":
		return s.Handle_ReadUser(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) Handle_CreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserReq := new(types.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&createUserReq); err != nil {
		return err
	}

	user := types.NewUser(
		createUserReq.Username,
		createUserReq.Password,
	)

	if err := s.store.Query_CreateUser(user); err != nil {
		return err
	}

	tokenString, err := utils.CreateJWT(user)
	if err != nil {
		return err
	}

	fmt.Println("JWT Token: ", tokenString)

	return utils.WriteJSON(w, http.StatusOK, user)
}

func (s *Server) Handle_ReadUser(w http.ResponseWriter, r *http.Request) error {
	users, err := s.store.Query_ReadUsers()
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, users)
}

func (s *Server) Handle_ReadUserById(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		id, err := utils.GetID(r)
		if err != nil {
			return err
		}

		users, err := s.store.Query_ReadUserByID(id)
		if err != nil {
			return err
		}

		return utils.WriteJSON(w, http.StatusOK, users)
	case "PUT":
		return s.Handle_UpdateUser(w, r)
	case "DELETE":
		return s.Handle_DeleteUser(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) Handle_UpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) Handle_DeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := utils.GetID(r)
	if err != nil {
		return err
	}

	if err := s.store.Query_DeleteUser(id); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]uuid.UUID{"deleted": id})
}
