package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	tableName struct{}   `sql:"password"`
	ID        string     `json:"id"`
	UserID    string     `json:"userId"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
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
