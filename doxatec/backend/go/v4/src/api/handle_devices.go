package api

import (
	"doxatec/types"
	"doxatec/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) Handle_Device(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.Handle_CreateDevice(w, r)
	case "GET":
		return s.Handle_ReadDevices(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) Handle_CreateDevice(w http.ResponseWriter, r *http.Request) error {
	createDeviceReq := new(types.CreateDeviceRequest)
	if err := json.NewDecoder(r.Body).Decode(&createDeviceReq); err != nil {
		return err
	}

	device := types.NewDevice(
		createDeviceReq.Name,
		createDeviceReq.Owner,
	)

	if err := s.store.Query_CreateDevice(device); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, device)
}

func (s *Server) Handle_ReadDevices(w http.ResponseWriter, r *http.Request) error {
	devices, err := s.store.Query_ReadDevices()
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, devices)
}

func (s *Server) Handle_ReadDeviceById(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		id, err := utils.GetID(r)
		if err != nil {
			return err
		}

		devices, err := s.store.Query_ReadDeviceByID(id)
		if err != nil {
			return err
		}

		return utils.WriteJSON(w, http.StatusOK, devices)
	case "PUT":
		return s.Handle_UpdateDevice(w, r)
	case "DELETE":
		return s.Handle_DeleteDevice(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *Server) Handle_UpdateDevice(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) Handle_DeleteDevice(w http.ResponseWriter, r *http.Request) error {
	id, err := utils.GetID(r)
	if err != nil {
		return err
	}

	if err := s.store.Query_DeleteDevice(id); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]uuid.UUID{"deleted": id})
}
