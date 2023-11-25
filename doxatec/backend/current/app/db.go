package app

import (
	"database/sql"
	"doxapi/utils"
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
		db_user     string = "gosql_l0kb_user"
		db_password string = "lPRWwZc2NJ630lgZolV4BPzLIO1cFFm8"
		db_host     string = "dpg-clgun9r1hq4c73bkh290-a.oregon-postgres.render.com"
		db_port     string = "5432"
		db_postgres string = "gosql_l0kb"
	)

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=verify-full", db_user, db_password, db_host, db_port, db_postgres))
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Minute * 3)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Database")

	return &Database{
		db: db,
	}, nil
}

func (s *Database) InitializeTables() {
	if _, err := s.db.Exec(utils.StringQL("sqls/clients/table/create.sql")); err != nil {
		log.Println(err)
	}
}
