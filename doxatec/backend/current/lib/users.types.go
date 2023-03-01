package lib

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func NewUser(username, password, avatar, name, email, phone string) *User {
	return &User{
		ID:       uuid.New(),
		Username: username,
		Password: password,
		Avatar:   avatar,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Created:  time.Now().UTC(),
		Modified: time.Now().UTC(),
	}
}
