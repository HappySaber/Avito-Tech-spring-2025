package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Printf("Error comparing passwords: %s", err)
	}
	return err == nil
}
