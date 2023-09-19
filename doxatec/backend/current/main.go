package main

import (
	"doxapi/app"
	"log"
)

func main() {
	db, err := app.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	db.InitializeTables()

	server := app.NewApi(":3000", db)
	server.Start()
}
