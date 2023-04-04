package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	port   string
	storer *Postgres
}

func NewApi(listenAddress string, store *Postgres) *Api {
	return &Api{
		port:   listenAddress,
		storer: store,
	}
}

func (s *Api) Start() {
	whitelist := map[string]bool{
		"":                       true,
		"http://localhost":       true,
		"http://0.0.0.0":         true,
		"http://192.168.0.1":     true,
		"http://172.17.0.1":      true,
		"https://142.93.207.120": true,
	}

	cors := func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			if whitelist[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Encoding, Authorization, Content-Type, Content-Length, Context-Type, X-CSRF-Token")
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Context-Type", "application/x-www-form-urlencoded")

				handler.ServeHTTP(w, r)
				return
			}

			http.Error(w, "origin not allowed", http.StatusForbidden)
		})
	}

	router := mux.NewRouter().StrictSlash(true)

	router.Use(cors)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello from Doxapi")
	})

	router.HandleFunc("/api/clients/list", s.ListClients).Methods("GET")
	router.HandleFunc("/api/clients/create", s.CreateClient).Methods("POST")
	router.HandleFunc("/api/clients/read/{id}", s.ReadClient).Methods("GET")
	router.HandleFunc("/api/clients/update/{id}", s.UpdateClient).Methods("PUT", "PATCH")
	router.HandleFunc("/api/clients/delete/{id}", s.DeleteClient).Methods("DELETE")

	log.Println("Doxapi available at port:", s.port)
	log.Fatal(http.ListenAndServe(s.port, router))
}
