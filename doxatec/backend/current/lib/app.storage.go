package lib

import "github.com/google/uuid"

type Storage interface {
	users_init() error
	users_beforeInsert(u *User) (*User, error)
	users_insert(u *User) error
	users_readTable() ([]*User, error)
	users_read(id uuid.UUID) (*User, error)
	users_readCol(column string, value any) error
	users_update(id uuid.UUID, column string, value any) error
	users_delete(id uuid.UUID) error
	Users_beforeSignin(username, password string) (*User, error)
}
