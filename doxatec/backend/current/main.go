package main

import (
	"doxatec/lib"
	"log"
)

func main() {
	store, err := lib.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := lib.NewServer(":80", store)
	server.Start()
}
