package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	const (
		db_user     string = "doxadmin"
		db_password string = "d@x@dm1n"
		db_host     string = "172.17.0.1" //"142.93.207.120"
		db_port     string = "5432"
		db_database string = "doxatec"
	)

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db_user, db_password, db_host, db_port, db_database))
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Minute * 3)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Database{
		db: db,
	}, nil
}

func (db *Database) Init() error {
	return nil
}
