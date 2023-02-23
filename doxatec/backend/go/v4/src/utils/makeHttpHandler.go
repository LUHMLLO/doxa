package utils

import (
	"doxatec/types"
	"net/http"
)

func MakeHttpHandleFunc(af types.ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := af(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, types.ApiError{Error: err.Error()})
		}
	}
}
