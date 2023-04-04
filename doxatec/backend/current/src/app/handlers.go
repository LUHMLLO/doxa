package app

import (
	"doxapi/utils"
	"encoding/json"
	"net/http"
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

// func (s *Api) ReadClient(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	w.Write([]byte(fmt.Sprintf("read client '%v'", vars["id"])))
// }

// func (s *Api) UpdateClient(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	w.Write([]byte(fmt.Sprintf("update/patch client '%v'", vars["id"])))
// }

// func (s *Api) DeleteClient(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	w.Write([]byte(fmt.Sprintf("delete client '%v'", vars["id"])))
// }
