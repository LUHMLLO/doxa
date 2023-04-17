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
	storer *Postgres
}

func NewApi(listenAddress string, store *Postgres) *Api {
	return &Api{
		port:   listenAddress,
		storer: store,
	}
}

func (s *Api) NewApiRoute(r *mux.Router, t interface{}, entity, path, method string) {
	endpoint := fmt.Sprintf("/api/%s/%s", entity, path)
	//log.Printf("router.HandleFunc(%s, s.HandlerList(%s, reflect.TypeOf(t))).Methods(%s)\n", endpoint, entity, method)
	r.HandleFunc(endpoint, s.HandlerList(entity, reflect.TypeOf(t))).Methods(method)
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

	s.NewApiRoute(router, Client{}, "clients", "list", "GET")
	s.NewApiRoute(router, NewClient{}, "clients", "create/{id}", "POST")
	s.NewApiRoute(router, Client{}, "clients", "read/{id}", "GET")
	s.NewApiRoute(router, UpdateClient{}, "clients", "update/{id}", "PATCH")
	s.NewApiRoute(router, Client{}, "clients", "delete/{id}", "DELETE")

	s.NewApiRoute(router, User{}, "users", "list", "GET")
	s.NewApiRoute(router, NewUser{}, "users", "create/{id}", "POST")
	s.NewApiRoute(router, User{}, "users", "read/{id}", "GET")
	s.NewApiRoute(router, UpdateUser{}, "users", "update/{id}", "PATCH")
	s.NewApiRoute(router, User{}, "users", "delete/{id}", "DELETE")

	s.NewApiRoute(router, Device{}, "devices", "list", "GET")
	s.NewApiRoute(router, NewDevice{}, "devices", "create/{id}", "POST")
	s.NewApiRoute(router, Device{}, "devices", "read/{id}", "GET")
	s.NewApiRoute(router, UpdateDevice{}, "devices", "update/{id}", "PATCH")
	s.NewApiRoute(router, Device{}, "devices", "delete/{id}", "DELETE")

	log.Println("Doxapi available at port:", s.port)
	log.Fatal(http.ListenAndServe(s.port, router))
}
