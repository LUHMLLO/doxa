package utils

import (
	"fmt"
	"log"
	"net/http"
)

func SetHeader(w http.ResponseWriter, r *http.Request, methods string) {
	origin := r.Header.Get("Origin")
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	if methods != "" {
		w.Header().Set("Access-Control-Allow-Methods", fmt.Sprintf("%s, OPTIONS", methods))
	} else {
		log.Fatal("need to set at least one method")
	}

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Context-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
}
