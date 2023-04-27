package app

import (
	"doxapi/utils"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func (s *Api) HandlerList(entity string, t reflect.Type) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entities, err := s.storer.QueryList(entity, t)

		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(entities)
	}
}

func (s *Api) HandlerCreate(entity string, requestBody interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		_, err = utils.ReorderMap(requestBody.(map[string]interface{}), reflect.TypeOf(requestBody).Elem())
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode("done")
	}
}

func (s *Api) HandlerRead(entity string, t reflect.Type) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		item, err := s.storer.QueryRead(entity, t, id)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(item)
	}
}

func (s *Api) HandlerUpdate(entity string, requestBody interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		item, err := s.storer.QueryUpdate(entity, id, requestBody)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(item)
	}
}

func (s *Api) HandlerDelete(entity string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		s.storer.QueryDelete(entity, id)

		json.NewEncoder(w).Encode("client deleted succesfully")
	}
}
