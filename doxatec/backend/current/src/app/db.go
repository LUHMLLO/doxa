package app

import (
	"database/sql"
	"doxapi/utils"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() (*Postgres, error) {
	const (
		db_user     string = "doxadmin"
		db_password string = "d@x@dm1n"
		db_host     string = "172.17.0.1" //"142.93.207.120"
		db_port     string = "5432"
		db_postgres string = "doxatec"
	)

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db_user, db_password, db_host, db_port, db_postgres))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Minute * 3)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Postgres{
		db: db,
	}, nil
}

func (s *Postgres) Init() error {
	utils.RowsQL(s.db, "sqls/clients/table/create.sql")
	utils.RowsQL(s.db, "sqls/users/table/create.sql")

	utils.RowsQL(s.db, "sqls/devices/table/create.sql")
	utils.RowsQL(s.db, "sqls/temperatures/table/create.sql")

	utils.RowsQL(s.db, "sqls/subscriptions/table/create.sql")
	utils.RowsQL(s.db, "sqls/transfers/table/create.sql")
	return nil
}
