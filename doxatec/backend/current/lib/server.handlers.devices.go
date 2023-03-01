package lib

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) Devices_Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		s.handle_allDevices(w, r)
	case "/create":
		s.handle_insertDevices(w, r)
	case "/read/{id}":
		s.handle_readDevices(w, r)
	case "/update/{id}":
		s.handle_updateDevices(w, r)
	case "/delete/{id}":
		s.handle_deleteDevices(w, r)
	}
}

func (s *Server) handle_allDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	devices, err := s.store.Query_allDevices()

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(devices)
}

func (s *Server) handle_insertDevices(w http.ResponseWriter, r *http.Request) {
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

	err = s.store.Query_insertDevices(device)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(device)
}

func (s *Server) handle_readDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	device, err := s.store.Query_readDevices(id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(device)
}

func (s *Server) handle_updateDevices(w http.ResponseWriter, r *http.Request) {
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

	err = s.store.Query_updateDevices(id, device)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"updated": id})
}

func (s *Server) handle_deleteDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.store.Query_deleteDevices(id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"deleted": id})
}
