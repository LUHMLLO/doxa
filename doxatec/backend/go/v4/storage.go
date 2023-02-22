package main

import (
	"database/sql"

	"github.com/google/uuid"
)

type Storage interface {
	CreateUser(*User) error
	ReadUser(uuid.UUID) (*User, error)
	UpdateUser(*User) error
	DeleteUser(uuid.UUID) error
}

type MySqlStore struct {
	db *sql.DB
}

func NewMySqlStore() (*MySqlStore, error) {

}
