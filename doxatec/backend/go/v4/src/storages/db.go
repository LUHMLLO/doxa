package storages

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connectionString := "user=postgres dbname=postgres password=d@x@dm1n sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	s.Query_CreateUserTable()
	s.Query_CreateProfileTable()
	s.Query_CreateDeviceTable()

	return nil
}
