package storage

import (
	"MyGoSql/types"

	"github.com/google/uuid"
)

type Storage interface {
	GetUser(uuid.UUID) *types.User
	GetOwner(uuid.UUID) *types.Owner
	GetDevice(uuid.UUID) *types.Device
}
