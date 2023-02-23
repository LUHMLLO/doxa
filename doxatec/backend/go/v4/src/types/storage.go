package types

import (
	"github.com/google/uuid"
)

type Storage interface {
	Query_CreateUser(*User) error
	Query_ReadUsers() ([]*User, error)
	Query_ReadUserByID(uuid.UUID) (*User, error)
	Query_UpdateUser(*User) error
	Query_DeleteUser(uuid.UUID) error
}
