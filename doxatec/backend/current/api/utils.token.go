package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const (
	SuperSecret string = "There is no secret actually"
)

func GenerateToken(id uuid.UUID) string {
	claims := jwt.MapClaims{
		"authorized": true,
		"id":         id,
		"expires":    time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(SuperSecret + RandSalt16()))
	if err != nil {
		return fmt.Sprintf("error generating token: %s", err.Error())
	}

	return signed
}
