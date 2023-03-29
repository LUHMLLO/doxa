package main

import (
	"log"
	"net/http"
)

func UsesToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			headerToken := r.Header.Get("Authorization")
			if headerToken == "" {
				log.Println("authorization does not exist's")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
}
