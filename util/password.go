package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (*string, error) {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	result := string(passwordHash)

	return &result, nil
}

func ComparePassword(hashedPassword, password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(hashedPassword)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
