package api

import (
	"doxatec/types"
	"doxatec/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) Handle_Profile(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.Handle_CreateProfile(w, r)
	case "GET":
		return s.Handle_ReadProfiles(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) Handle_CreateProfile(w http.ResponseWriter, r *http.Request) error {
	createProfileReq := new(types.CreateProfileRequest)
	if err := json.NewDecoder(r.Body).Decode(&createProfileReq); err != nil {
		return err
	}

	profile := types.NewProfile(
		createProfileReq.Avatar,
		createProfileReq.Name,
		createProfileReq.Email,
		createProfileReq.Phone,
	)

	if err := s.store.Query_CreateProfile(profile); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, profile)
}

func (s *Server) Handle_ReadProfiles(w http.ResponseWriter, r *http.Request) error {
	profiles, err := s.store.Query_ReadProfiles()
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, profiles)
}

func (s *Server) Handle_ReadProfileById(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		id, err := utils.GetID(r)
		if err != nil {
			return err
		}

		profiles, err := s.store.Query_ReadProfileByID(id)
		if err != nil {
			return err
		}

		return utils.WriteJSON(w, http.StatusOK, profiles)
	case "PUT":
		return s.Handle_UpdateProfile(w, r)
	case "DELETE":
		return s.Handle_DeleteProfile(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) Handle_UpdateProfile(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) Handle_DeleteProfile(w http.ResponseWriter, r *http.Request) error {
	id, err := utils.GetID(r)
	if err != nil {
		return err
	}

	if err := s.store.Query_DeleteProfile(id); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]uuid.UUID{"deleted": id})
}
