package lib

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const secretJWTkey string = "doxasecret"

func GenerateJWT(id uuid.UUID) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    id.String(),
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(), //12 hours
	})

	token, err := claims.SignedString([]byte(secretJWTkey))
	if err != nil {
		return "", fmt.Errorf("error generating token: %s", err.Error())
	}

	return token, nil
}

type httpClaimsContext struct {
	Issuer    string
	ExpiresAt int64
	Status    bool
}

var ClaimsContext = &httpClaimsContext{}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("jwt")
			if err != nil {
				// set error status code and message
				http.Error(w, "token does not exists", 200)
				return
			}

			if cookie.Value == "" {
				// set error status code and message
				http.Error(w, "received an empty token", 400)
				return
			}

			token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretJWTkey), nil
			})
			if err != nil {
				// set error status code and message
				http.Error(w, fmt.Sprintf("not authenticated: %v", err), 400)
				return
			}

			claims, ok := token.Claims.(*jwt.StandardClaims)
			// log.Println("claims: ", claims)
			// log.Println("status: ", ok)

			r = r.WithContext(
				context.WithValue(
					r.Context(),
					ClaimsContext,
					httpClaimsContext{
						Issuer:    claims.Issuer,
						ExpiresAt: claims.ExpiresAt,
						Status:    ok,
					},
				),
			)

			next.ServeHTTP(w, r)
		})
}
