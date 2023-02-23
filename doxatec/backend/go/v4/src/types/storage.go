package types

import (
	"github.com/google/uuid"
)

type Storage interface {
	CreateUser(*User) error
	ReadUsers() ([]*User, error)
	ReadUserByID(uuid.UUID) (*User, error)
	UpdateUser(*User) error
	DeleteUser(uuid.UUID) error
}
