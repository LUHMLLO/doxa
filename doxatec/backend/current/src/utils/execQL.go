package utils

import (
	"database/sql"
	"log"
	"os"
)

func ExecQL(db *sql.DB, path string) sql.Result {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec(string(bytes))
	if err != nil {
		log.Fatal(err)
	}

	return result
}
