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

	Query_CreateProfile(*Profile) error
	Query_ReadProfiles() ([]*Profile, error)
	Query_ReadProfileByID(uuid.UUID) (*Profile, error)
	Query_UpdateProfile(*Profile) error
	Query_DeleteProfile(uuid.UUID) error

	Query_CreateDevice(*Device) error
	Query_ReadDevices() ([]*Device, error)
	Query_ReadDeviceByID(uuid.UUID) (*Device, error)
	Query_UpdateDevice(*Device) error
	Query_DeleteDevice(uuid.UUID) error
}
