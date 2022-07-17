package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	tableName struct{} `pg:"password"`
	ID        string
	UserID    string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}

func (p *Password) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Password = string(passwordHash)

	return nil
}

func (p *Password) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(p.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
