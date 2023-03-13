package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Profile  Profile   `json:"profile"`
	Time     Stamp     `json:"time_stamp"`
}

type UserRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Profile struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}

type ProfileRequestBody struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`
}

type Device struct {
	ID    uuid.UUID `json:"id"`
	Owner string    `json:"owner"`
	Name  string    `json:"name"`
	PIN   string    `json:"pin"`
	Temps Temps     `json:"temps"`
	Time  Stamp     `json:"time_stamp"`
}

type DeviceRequestBody struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type Temps struct {
	TempSup float64 `json:"temp_sup"`
	TempMid float64 `json:"temp_mid"`
	TempSub float64 `json:"temp_sub"`
}

type Stamp struct {
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Seen     time.Time `json:"seen"`
}

func NewUser(username, password, avatar, name, email, phone, role string) *User {
	return &User{
		ID:       uuid.New(),
		Username: username,
		Password: password,
		Role:     role,

		Profile: Profile{
			Avatar: avatar,
			Name:   name,
			Email:  email,
			Phone:  phone,
		},

		Time: Stamp{
			Created:  time.Now().UTC(),
			Modified: time.Now().UTC(),
			Seen:     time.Now().UTC(),
		},
	}
}
