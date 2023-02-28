package lib

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Storage interface {
	Users_CreateTable() error
	Users_InsertToTable(u *User) error
	Users_ReadFromTable() ([]*User, error)
	Users_ReadFromTableByID(id uuid.UUID) (*User, error)
	Users_UpdateFromTableByID(id uuid.UUID, u *User) error
	Users_DeleteFromTableByID(id uuid.UUID) (uuid.UUID, error)
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
	s.Users_CreateTable()

	return nil
}
