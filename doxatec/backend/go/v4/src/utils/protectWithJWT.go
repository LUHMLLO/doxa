package utils

import (
	"doxatec/types"
	"fmt"
	"net/http"
	"os"
	"reflect"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateJWT(user *types.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := &jwt.MapClaims{
		"ExpiresAt": 15000,
		"UserID":    user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ProtectWithJWT(handlerFunc http.HandlerFunc, store types.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Calling JWT auth middleware")

		tokenString := r.Header.Get("x-jwt-token")
		token, err := ValidateJWT(tokenString)
		if err != nil {
			PermissionDenied(w)
			return
		}
		if !token.Valid {
			PermissionDenied(w)
			return
		}

		userID, err := GetID(r)
		if err != nil {
			PermissionDenied(w)
			return
		}

		user, err := store.Query_ReadUserByID(userID)
		if err != nil {
			PermissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(reflect.TypeOf(claims["UserID"]))
		if user.ID.String() != string(claims["UserID"].(string)) {
			PermissionDenied(w)
		}

		if err != nil {
			PermissionDenied(w)
			return
		}

		handlerFunc(w, r)
	}
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}
