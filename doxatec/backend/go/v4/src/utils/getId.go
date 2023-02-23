package utils

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetID(r *http.Request) (uuid.UUID, error) {
	idString := mux.Vars(r)["id"]
	id, err := uuid.Parse(idString)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idString)
	}
	return id, nil
}
