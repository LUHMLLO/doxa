package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Client   uuid.UUID `json:"client_id"`
	Avatar   string    `json:"avatar"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Accessed time.Time `json:"accessed"`
}
type NewUser struct {
	Client   uuid.UUID `json:"client_id"`
	Avatar   string    `json:"avatar"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
}
