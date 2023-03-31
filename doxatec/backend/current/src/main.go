package main

import "log"

func main() {
	db, err := NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewServer(":3000", db)
	server.Start()
}
