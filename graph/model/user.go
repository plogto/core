package model

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	tableName struct{} `pg:"user"`
	ID        string
	Username  string
	Email     string
	FullName  string
	Bio       *string
	Role      string
	IsPrivate bool `pg:",use_zero"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}

func (u *User) GenToken() (*AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        u.ID,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "plog",
	})

	token, err := t.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		Token:     token,
		ExpiredAt: expiredAt,
	}, nil
}
