package lib

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	listenAddress string
	store         Storage
}

func NewServer(listenAddress string, store Storage) *Server {
	return &Server{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *Server) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/create/user", s.CreateUser)
	router.HandleFunc("/read/users", s.ReadUsers)

	log.Println("Doxatec server running on port:", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := []interface{}{
		uuid.New(),
		"@Democlient",
		"1234",
		"https://cdn.dribbble.com/userupload/4987411/file/original-c317d418c0b2b734856b4b8c8db4370e.jpg?compress=1&resize=752x",
		"Demo Client",
		"client@demo.com",
		"(809)-555-1234",
		time.Now().UTC(),
		time.Now().UTC(),
	}
	s.store.InsertToTable("users", user)
}

func (s *Server) ReadUsers(w http.ResponseWriter, r *http.Request) {
	s.store.ReadFromTable("users")
}
