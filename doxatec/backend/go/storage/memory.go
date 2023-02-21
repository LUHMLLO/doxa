package storage

import (
	"MyGoSql/types"
	"time"

	"github.com/google/uuid"
)

type MemoryStorage struct {
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) GetUser(uuid.UUID) *types.User {
	return &types.User{
		ID:       uuid.New(),
		Username: "@Dev",
		Password: "1234567890",
	}
}

func (s *MemoryStorage) GetOwner(uuid.UUID) *types.Owner {
	return &types.Owner{
		ID:    uuid.New(),
		Name:  "Client Dev",
		Phone: "(809)-129-6767",
		Email: "client@dev.com",
	}
}

func (s *MemoryStorage) GetDevice(uuid.UUID) *types.Device {
	return &types.Device{
		ID:      uuid.New(),
		Name:    "Client Dev - Fridge #1",
		Owner:   "DevClient",
		TempSup: 10,
		TempMid: 5,
		TempSub: 0,
		Date:    time.Now(),
	}
}
