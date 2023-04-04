package app

import (
	"doxapi/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

func (s *Api) ListClients(w http.ResponseWriter, r *http.Request) {
	rows := utils.RowsQL(s.storer.db, "sqls/clients/table/read.sql")

	clients := []*Client{}

	for rows.Next() {
		client := &Client{}

		if err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.Email,
			&client.Phone,
			&client.Created,
			&client.Modified,
		); err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		clients = append(clients, client)
	}

	json.NewEncoder(w).Encode(clients)
}

func (s *Api) CreateClient(w http.ResponseWriter, r *http.Request) {
	DTO := &NewClient{}

	err := json.NewDecoder(r.Body).Decode(&DTO)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	_, err = s.storer.db.Query(utils.StringQL("sqls/clients/crud/create.sql"), DTO.Name, DTO.Email, DTO.Phone)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (s *Api) ReadClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	rows, err := s.storer.db.Query(utils.StringQL("sqls/clients/crud/read.sql"), id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	client := &Client{}

	for rows.Next() {
		if err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.Email,
			&client.Phone,
			&client.Created,
			&client.Modified,
		); err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
	}

	json.NewEncoder(w).Encode(client)
}

func (s *Api) UpdateClient(w http.ResponseWriter, r *http.Request) {
	DTO := &UpdateClient{}
	err := json.NewDecoder(r.Body).Decode(&DTO)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if DTO.Name != "" {
		query := fmt.Sprintf(utils.StringQL("sqls/clients/crud/update.sql"), pq.QuoteIdentifier("name"))

		_, err := s.storer.db.Query(query, id, DTO.Name)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
	}

	if DTO.Email != "" {
		query := fmt.Sprintf(utils.StringQL("sqls/clients/crud/update.sql"), pq.QuoteIdentifier("email"))

		_, err := s.storer.db.Query(query, id, DTO.Email)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
	}

	if DTO.Phone != "" {
		query := fmt.Sprintf(utils.StringQL("sqls/clients/crud/update.sql"), pq.QuoteIdentifier("phone"))

		_, err := s.storer.db.Query(query, id, DTO.Phone)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
	}

	_, err = s.storer.db.Query(fmt.Sprintf(utils.StringQL("sqls/clients/crud/update.sql"), pq.QuoteIdentifier("modified")), id, time.Now().UTC())
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	rows, err := s.storer.db.Query(utils.StringQL("sqls/clients/crud/read.sql"), id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	client := &Client{}

	for rows.Next() {
		if err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.Email,
			&client.Phone,
			&client.Created,
			&client.Modified,
		); err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
	}

	json.NewEncoder(w).Encode(client)
}

func (s *Api) DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := s.storer.db.Query(utils.StringQL("sqls/clients/crud/delete.sql"), id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}
