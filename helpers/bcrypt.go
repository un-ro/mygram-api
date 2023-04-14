package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(pass string) (hashed string) {
	salt := 8
	password := []byte(pass)

	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		log.Fatal("error generate password")
		return
	}

	return string(hash)
}

func ComparePassword(hashed, pass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, pass)
	if err != nil {
		log.Fatal("Error compare hash and pass")
		return false
	}

	return true
}
