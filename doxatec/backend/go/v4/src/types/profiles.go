package types

import (
	"time"

	"github.com/google/uuid"
)

type CreateProfileRequest struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}

type Profile struct {
	ID     uuid.UUID `json:"id"`
	Avatar string    `json:"avatar"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Phone  string    `json:"phone"`

	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func NewProfile(avatar, name, email, phone string) *Profile {
	return &Profile{
		ID:     uuid.New(),
		Avatar: avatar,
		Name:   name,
		Email:  email,
		Phone:  phone,

		Created:  time.Now().UTC(),
		Modified: time.Now().UTC(),
	}
}
