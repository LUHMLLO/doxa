package api

import (
	"doxatec/types"
	"doxatec/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddress string
	store         types.Storage
}

func NewApiServer(listenAddress string, store types.Storage) *ApiServer {
	return &ApiServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *ApiServer) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/users", utils.MakeHttpHandleFunc(s.handleUser))
	router.HandleFunc("/users/{id}", utils.MakeHttpHandleFunc(s.handleReadUserById))

	log.Println("DOXA api server ruuning on port:", s.listenAddress)
	log.Printf("http://localhost%s\n", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)
}
