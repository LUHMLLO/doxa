package main

import (
	"database/sql"
	"log"
	"os"
)

func ExecQL(db *sql.DB, path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(bytes))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Queried")
}
