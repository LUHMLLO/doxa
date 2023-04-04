package utils

import (
	"database/sql"
	"log"
	"os"
)

func StringQL(path string) string {
	sql, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	return string(sql)
}

func RowsQL(db *sql.DB, path string) *sql.Rows {
	sql, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	rows, err := db.Query(string(sql))
	if err != nil {
		log.Println(err)
	}

	return rows
}
