package main

import (
	"crypto/rand"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func RandSalt16() string {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		log.Fatal(err)
	}
	return string(salt)
}

func SecurePassword(pass, salt string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass+salt), 14)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func ValidatePassword(pass, salt, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass+salt))
	return err == nil
}
