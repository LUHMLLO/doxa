package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Avatar   string    `json:"avatar"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Customer string    `json:"customer"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Accessed time.Time `json:"accessed"`
}

func NewUser(avatar, username, password string) *User {
	return &User{
		ID:       uuid.New(),
		Avatar:   avatar,
		Username: username,
		Password: password,
		Customer: uuid.NewString(),
		Created:  time.Now(),
		Modified: time.Now(),
		Accessed: time.Now(),
	}
}
