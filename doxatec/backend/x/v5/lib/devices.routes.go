package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) Route_all_devices(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, r, "GET")

	devices, err := s.store.devices_readTable()

	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading table: %v", err))
		return
	}

	json.NewEncoder(w).Encode(devices)
}

func (s *Server) Route_insert_device(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, r, "POST")

	createDeviceReq := &CreateDeviceRequest{}
	err := json.NewDecoder(r.Body).Decode(&createDeviceReq)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading request body: %v", err))
		return
	}

	device := NewDevice(
		createDeviceReq.Owner,
		createDeviceReq.Name,
		"",
		0,
		0,
		0,
	)

	_, err = s.store.devices_beforeInsert(device)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error before insertion: %v", err))
		return
	}

	err = s.store.devices_insert(device)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error during insertion: %v", err))
		return
	}

	json.NewEncoder(w).Encode(device)
}

func (s *Server) Route_read_device(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, r, "GET")

	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading route params: %v", err))
		return
	}

	device, err := s.store.devices_read(id)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading device: %v", err))
		return
	}

	json.NewEncoder(w).Encode(device)
}

func (s *Server) Route_update_device(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, r, "PUT")

	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading route params: %v", err))
		return
	}

	updateDeviceReq := &UpdateDeviceRequest{}
	err = json.NewDecoder(r.Body).Decode(&updateDeviceReq)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading request body: %v", err))
		return
	}

	device := NewDevice(
		updateDeviceReq.Owner,
		updateDeviceReq.Name,
		updateDeviceReq.PIN,
		updateDeviceReq.TempSup,
		updateDeviceReq.TempMid,
		updateDeviceReq.TempSub,
	)

	if device.Owner != "" {
		if err = s.store.devices_update(id, "owner", device.Owner); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating owner: %v", err))
			return
		}
	}

	if device.Name != "" {
		if err = s.store.devices_update(id, "name", device.Name); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating name: %v", err))
			return
		}
	}

	if device.PIN != "" {
		device.PIN, err = GeneratehashPassword(device.PIN)
		if err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error hashing password: %v", err))
			return
		}

		if err = s.store.devices_update(id, "pin", device.PIN); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating password: %v", err))
			return
		}
	}

	if device.TempSup != 0.0 {
		if err = s.store.devices_update(id, "tempsup", device.TempSup); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating temp sup: %v", err))
			return
		}
	}

	if device.TempMid != 0.0 {
		if err = s.store.devices_update(id, "tempmid", device.TempMid); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating temp mid: %v", err))
			return
		}
	}

	if device.TempSub != 0.0 {
		if err = s.store.devices_update(id, "tempsub", device.TempSub); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("error updating temp sub: %v", err))
			return
		}
	}

	if err = s.store.devices_update(id, "modified", time.Now().UTC()); err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error updating modified: %v", err))
		return
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"updated": id})
}

func (s *Server) Route_delete_device(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, r, "DELETE")

	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading route params: %v", err))
		return
	}

	err = s.store.devices_delete(id)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error deleting device: %v", err))
		return
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"deleted": id})
}
