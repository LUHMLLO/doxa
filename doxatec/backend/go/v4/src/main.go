package main

import (
	"doxatec/api"
	"doxatec/storages"
	"log"
)

func main() {
	store, err := storages.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(":3000", store)
	server.Start()
}
