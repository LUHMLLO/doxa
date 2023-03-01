package utils

import (
	"doxatec/types"
	"net/http"
)

func PermissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, types.ApiError{Error: "permission denied"})
}
