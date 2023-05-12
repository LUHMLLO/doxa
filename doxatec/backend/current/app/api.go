package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

type Api struct {
	port   string
	storer *Database
}

func NewApi(listenAddress string, store *Database) *Api {
	return &Api{
		port:   listenAddress,
		storer: store,
	}
}

func (s *Api) NewApiRoute(r *mux.Router, t interface{}, entity, path string) {
	endpoint := fmt.Sprintf("/api/%s/%s", entity, path)
	switch path {
	case "list":
		r.HandleFunc(endpoint, s.HandlerList(entity, reflect.TypeOf(t))).Methods("GET")
		return
	case "create":
		r.HandleFunc(endpoint, s.HandlerCreate(entity, t)).Methods("POST")
		return
	case "read/{id}":
		r.HandleFunc(endpoint, s.HandlerRead(entity, reflect.TypeOf(t))).Methods("GET")
		return
	case "update/{id}":
		r.HandleFunc(endpoint, s.HandlerUpdate(entity, reflect.TypeOf(t))).Methods("UPDATE")
		return
	case "delete/{id}":
		r.HandleFunc(endpoint, s.HandlerDelete(entity)).Methods("DELETE")
		return
	}
}

func (s *Api) Start() {
	whitelist := map[string]bool{
		"":                       true,
		"http://localhost":       true,
		"http://0.0.0.0":         true,
		"http://192.168.0.1":     true,
		"http://172.17.0.1":      true,
		"http://142.93.207.120":  true,
		"https://localhost":      true,
		"https://0.0.0.0":        true,
		"https://192.168.0.1":    true,
		"https://172.17.0.1":     true,
		"https://142.93.207.120": true,
	}

	router := mux.NewRouter().StrictSlash(true)

	router.Use(func(handler http.Handler) http.Handler {
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
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello from Doxapi")
	})

	s.NewApiRoute(router, Client{}, "clients", "list")
	s.NewApiRoute(router, NewClient{}, "clients", "create")
	s.NewApiRoute(router, Client{}, "clients", "read/{id}")
	s.NewApiRoute(router, UpdateClient{}, "clients", "update/{id}")
	s.NewApiRoute(router, Client{}, "clients", "delete/{id}")

	s.NewApiRoute(router, User{}, "users", "list")
	s.NewApiRoute(router, NewUser{}, "users", "create")
	s.NewApiRoute(router, User{}, "users", "read/{id}")
	s.NewApiRoute(router, UpdateUser{}, "users", "update/{id}")
	s.NewApiRoute(router, User{}, "users", "delete/{id}")

	s.NewApiRoute(router, Device{}, "devices", "list")
	s.NewApiRoute(router, NewDevice{}, "devices", "create")
	s.NewApiRoute(router, Device{}, "devices", "read/{id}")
	s.NewApiRoute(router, UpdateDevice{}, "devices", "update/{id}")
	s.NewApiRoute(router, Device{}, "devices", "delete/{id}")

	log.Println("Doxapi available at port:", s.port)
	log.Fatal(http.ListenAndServe(s.port, router))
}
