package lib

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) Route_all_devices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	devices, err := s.store.Query_read_all_devices_from_table()

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(devices)
}

func (s *Server) Route_insert_device(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	createDeviceReq := &CreateDeviceRequest{}
	err := json.NewDecoder(r.Body).Decode(&createDeviceReq)
	if err != nil {
		log.Fatalln("error in reading body: ", err.Error())
		return
	}

	device := NewDevice(
		createDeviceReq.Owner,
		createDeviceReq.Name,
	)

	_, err = s.store.Query_before_insert_device(device)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = s.store.Query_insert_device(device)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(device)
}

func (s *Server) Route_read_device(w http.ResponseWriter, r *http.Request) {
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

	device, err := s.store.Query_read_device_where_column("id", id.String())
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(device)
}

func (s *Server) Route_update_device(w http.ResponseWriter, r *http.Request) {
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

	createDeviceReq := &CreateDeviceRequest{}
	err = json.NewDecoder(r.Body).Decode(&createDeviceReq)
	if err != nil {
		log.Fatal(err)
	}

	device := NewDevice(
		createDeviceReq.Owner,
		createDeviceReq.Name,
	)

	if device.Name != "" {
		if err = s.store.Query_update_device_where_ID_and_column(id, "Name", device.Name); err != nil {
			log.Fatal(err)
		}
	}

	if err = s.store.Query_update_device_where_ID_and_column(id, "modified", time.Now().UTC()); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"updated": id})
}

func (s *Server) Route_delete_device(w http.ResponseWriter, r *http.Request) {
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

	_, err = s.store.Query_delete_device_where_ID(id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(map[string]uuid.UUID{"deleted": id})
}
