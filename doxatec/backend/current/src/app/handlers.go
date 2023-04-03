package app

import (
	entities "doxapi/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func (s *Api) ListClients(w http.ResponseWriter, r *http.Request) {
	sql, err := os.ReadFile("sqls/clients/table/read.sql")
	if err != nil {
		log.Println(err)
	}

	rows, err := s.storer.db.Query(string(sql))
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	clients := []*entities.Client{}

	for rows.Next() {
		client := &entities.Client{}

		if err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.Email,
			&client.Phone,
			&client.Created,
			&client.Modified,
		); err != nil {
			log.Println(err)
		}

		clients = append(clients, client)
	}

	json.NewEncoder(w).Encode(clients)
}

func (s *Api) CreateClient(w http.ResponseWriter, r *http.Request) {
	DTO := &entities.NewClient{}
	err := json.NewDecoder(r.Body).Decode(&DTO)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	sql, err := os.ReadFile("sqls/clients/crud/create.sql")
	if err != nil {
		log.Println(err)
		return
	}

	rows, err := s.storer.db.Query(string(sql), DTO.Name, DTO.Email, DTO.Phone)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	client := &entities.Client{}

	for rows.Next() {
		if err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.Email,
			&client.Phone,
			&client.Created,
			&client.Modified,
		); err != nil {
			log.Println(err)
			return
		}
	}

	json.NewEncoder(w).Encode(client)
}

func (s *Api) ReadClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Write([]byte(fmt.Sprintf("read client '%v'", vars["id"])))
}

func (s *Api) UpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Write([]byte(fmt.Sprintf("update/patch client '%v'", vars["id"])))
}

func (s *Api) DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Write([]byte(fmt.Sprintf("delete client '%v'", vars["id"])))
}
