package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashedPassword returns the bcrypt has of the password
func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error: for hashing password: %v", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if provided password is correct or not
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
