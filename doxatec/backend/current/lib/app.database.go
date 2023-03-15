package lib

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	const (
		db_user     string = "doxadmin"
		db_password string = "d@x@dm1n"
		db_host     string = "142.93.207.120"
		db_port     string = "5432"
		db_database string = "doxatec"
	)

	//url_local := "dbname=doxatec user=doxadmin password=d@x@dm1n sslmode=disable"
	url_public := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db_user, db_password, db_host, db_port, db_database)

	db, err := sql.Open("postgres", url_public)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (s *Database) Init() error {
	s.users_init()
	s.devices_init()

	return nil
}
