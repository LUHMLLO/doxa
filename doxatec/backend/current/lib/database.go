package lib

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Storage interface {
	Query_tableUsers() error
	Query_insertUsers(u *User) error
	Query_allUsers() ([]*User, error)
	Query_readUsers(id uuid.UUID) (*User, error)
	Query_updateUsers(id uuid.UUID, u *User) error
	Query_deleteUsers(id uuid.UUID) (uuid.UUID, error)

	Query_tableDevices() error
	Query_insertDevices(u *Device) error
	Query_allDevices() ([]*Device, error)
	Query_readDevices(id uuid.UUID) (*Device, error)
	Query_updateDevices(id uuid.UUID, u *Device) error
	Query_deleteDevices(id uuid.UUID) (uuid.UUID, error)
}

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "dbname=doxatec user=doxadmin password=d@x@dm1n sslmode=disable")
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
	s.Query_tableUsers()
	s.Query_tableDevices()

	return nil
}
