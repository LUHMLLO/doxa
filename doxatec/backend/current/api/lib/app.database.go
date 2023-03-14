package lib

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	// local db
	//db, err := sql.Open("postgres", "dbname=doxatec user=doxadmin password=d@x@dm1n sslmode=disable")

	// online db
	db, err := sql.Open("postgres", "postgres://doxadmin:DXIq5uCiqi2xQpUVokA55sRZHQPG4q32@dpg-cg8a6v9mbg53mc4sp150-a.oregon-postgres.render.com/doxatec")
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
