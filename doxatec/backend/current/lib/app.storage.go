package lib

import "github.com/google/uuid"

type Storage interface {
	users_init() error
	users_beforeInsert(u *User) (*User, error)
	users_insert(u *User) error
	users_readTable() ([]*User, error)
	users_read(id uuid.UUID) (*User, error)
	users_readCol(column string, value any) (*User, error)
	users_update(id uuid.UUID, column string, value any) error
	users_delete(id uuid.UUID) error
	Users_beforeSignin(username, password string) (*User, error)

	devices_init() error
	devices_beforeInsert(d *Device) (*Device, error)
	devices_insert(d *Device) error
	devices_readTable() ([]*Device, error)
	devices_read(id uuid.UUID) (*Device, error)
	devices_readCol(column string, value any) error
	devices_update(id uuid.UUID, column string, value any) error
	devices_delete(id uuid.UUID) error
}
