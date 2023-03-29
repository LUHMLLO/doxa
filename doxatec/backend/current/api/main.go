package main

import (
	"log"
)

func main() {
	password := "1234567890"
	salt := RandSalt16()

	SecuredPassword := SecurePassword(password, salt)
	log.Println("Secured: ", SecuredPassword)

	ValidatedPassword := ValidatePassword(password, salt, SecuredPassword)
	log.Println("Validated: ", ValidatedPassword)
}
