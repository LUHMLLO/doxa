package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	port string
	db   *Postgres
}

func NewApi(listenAddress string, database *Postgres) *Api {
	return &Api{
		port: listenAddress,
		db:   database,
	}
}

func (s *Api) Start() {
	router := mux.NewRouter()

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
