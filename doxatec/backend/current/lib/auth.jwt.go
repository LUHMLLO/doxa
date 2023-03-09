package lib

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userId string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    userId,
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(), //12 hours
	})

	token, err := claims.SignedString([]byte(secretJWTkey))
	if err != nil {
		return "", fmt.Errorf("error generating token: %s", err.Error())
	}

	return token, nil
}
