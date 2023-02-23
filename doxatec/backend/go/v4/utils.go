package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, variable any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(variable)
}

type ApiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func MakeHttpHandleFunc(af ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := af(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func scanIntoUser(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(&user.ID, &user.Avatar, &user.Username, &user.Password, &user.Customer, &user.Created, &user.Modified, &user.Accessed)

	return user, err
}

func getID(r *http.Request) (uuid.UUID, error) {
	idString := mux.Vars(r)["id"]
	id, err := uuid.Parse(idString)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idString)
	}
	return id, nil
}
