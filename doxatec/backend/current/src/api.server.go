package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port string
}

func NewStart(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Start() {
	router := mux.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("x-token", "Hello, world!")
			next.ServeHTTP(w, r)
		})
	})

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			origin := r.Header.Get("Origin")
			if origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			}

			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Context-Type", "application/x-www-form-urlencoded")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from root")
	})

	router.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from init")
	})

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from api")
	})

	log.Println("Doxapi available at port:", s.port)
	log.Fatal(http.ListenAndServe(s.port, router))
}
