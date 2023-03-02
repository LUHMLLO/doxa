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
	Role     string `json:"role"`
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Role     string    `json:"role"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func NewUser(username, password, avatar, name, email, phone, role string) *User {
	return &User{
		ID:       uuid.New(),
		Username: username,
		Password: password,
		Avatar:   avatar,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Role:     role,
		Created:  time.Now().UTC(),
		Modified: time.Now().UTC(),
	}
}

type SigninUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type SigninUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewSigninUser(username, password string) *SigninUser {
	return &SigninUser{
		Username: username,
		Password: password,
	}
}

type SecretJWTtoken struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	JWT      string `json:"jwt"`
}

func NewSecretJWTtoken(username, role, jwt string) *SecretJWTtoken {
	return &SecretJWTtoken{
		Username: username,
		Role:     role,
		JWT:      jwt,
	}
}
