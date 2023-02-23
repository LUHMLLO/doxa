package types

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`

	Profile string `json:"profile"`

	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Accessed time.Time `json:"accessed"`
}

func NewUser(username, password string) *User {
	return &User{
		ID:       uuid.New(),
		Username: username,
		Password: password,

		Profile: "",

		Created:  time.Now().UTC(),
		Modified: time.Now().UTC(),
		Accessed: time.Now().UTC(),
	}
}
