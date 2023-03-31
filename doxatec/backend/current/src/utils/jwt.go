package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func NewJWT(id uuid.UUID, role string) string {
	claims := jwt.MapClaims{
		"authorized": true,
		"id":         id,
		"role":       role,
		"created":    time.Now().Unix(),
		"expires":    time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(SuperSecret))
	if err != nil {
		log.Fatal("Error signing token: ", err)
	}

	return signed
}

func ValidateJWT(x_token string) {
	token, err := jwt.Parse(x_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return []byte(SuperSecret), nil
	})

	if err != nil {
		log.Println(err)
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Println("invalid token")
	}
}
