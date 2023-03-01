package lib

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) ReadAllDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	devices, err := s.store.Devices_ReadFromTable()

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(devices)
}

func (s *Server) InsertDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	createDeviceReq := &CreateDeviceRequest{}
	err := json.NewDecoder(r.Body).Decode(&createDeviceReq)
	if err != nil {
		log.Fatal(err)
	}

	device := NewDevice(
		createDeviceReq.Owner,
		createDeviceReq.Name,
		createDeviceReq.TempSup,
		createDeviceReq.TempMid,
		createDeviceReq.TempSub,
	)

	err = s.store.Devices_InsertToTable(device)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(device)
}

func (s *Server) ReadDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	device, err := s.store.Devices_ReadFromTableByID(id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(device)
}

func (s *Server) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	createDeviceReq := &CreateDeviceRequest{}
	err = json.NewDecoder(r.Body).Decode(&createDeviceReq)
	if err != nil {
		log.Fatal(err)
	}

	device := NewDevice(
		createDeviceReq.Owner,
		createDeviceReq.Name,
		createDeviceReq.TempSup,
		createDeviceReq.TempMid,
		createDeviceReq.TempSub,
	)

	err = s.store.Devices_UpdateFromTableByID(id, device)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"updated": id})
}

func (s *Server) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.store.Devices_DeleteFromTableByID(id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"deleted": id})
}
