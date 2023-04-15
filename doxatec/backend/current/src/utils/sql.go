package utils

import (
	"database/sql"
	"log"
	"os"
)

func StringQL(path string) string {
	sql, err := os.ReadFile(path)
	if err != nil {
		log.Println("StringQL error: ", err)
	}

	return string(sql)
}

func RowsQL(db *sql.DB, path string, params ...interface{}) *sql.Rows {
	sql, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	var args []interface{}
	for _, arg := range params {
		if m, ok := arg.(map[string]interface{}); ok {
			for _, key := range sortedKeys(m) {
				args = append(args, m[key])
			}
		} else {
			args = append(args, arg)
		}
	}

	rows, err := db.Query(string(sql), args...)
	if err != nil {
		log.Println("RowsQL error: ", err)
	}

	return rows
}

func ExecQL(db *sql.DB, path string, params ...interface{}) sql.Result {
	sql, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	var args []interface{}
	for _, arg := range params {
		if m, ok := arg.(map[string]interface{}); ok {
			for _, key := range sortedKeys(m) {
				args = append(args, m[key])
			}
		} else {
			args = append(args, arg)
		}
	}

	result, err := db.Exec(string(sql), args...)
	if err != nil {
		log.Println("ExecQL error: ", err)
	}

	return result
}
