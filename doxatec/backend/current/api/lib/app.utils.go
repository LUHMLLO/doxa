package lib

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func StringToQuery(slice []string) string {
	lastKey := slice[len(slice)-1]

	var keys strings.Builder

	for _, value := range slice {
		if value == lastKey {
			keys.WriteString(fmt.Sprintf("%s\n", value))
		} else {
			keys.WriteString(fmt.Sprintf("%s,\n", value))
		}
	}

	return keys.String()
}

func SetHeaders(w http.ResponseWriter, r *http.Request, methods string) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Credentials, Headers, Methods, Origin, Content-Type, Context-Type")
	w.Header().Set("Access-Control-Allow-Methods", methods)

	origin := r.Header.Get("Origin")
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
