package main

import "golang.org/x/crypto/bcrypt"

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
