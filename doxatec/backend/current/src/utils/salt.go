package main

import (
	"crypto/rand"
	"log"
)

func NewRandSalt16() string {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		log.Fatal(err)
	}
	return string(salt)
}
