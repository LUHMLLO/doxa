package main

import (
	"doxapi/app"
	"log"
)

func main() {
	db, err := app.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	server := app.NewApi(":3000", db)
	server.Start()
}
