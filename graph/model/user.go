package model

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	tableName       struct{} `pg:"users"`
	ID              string
	Username        string
	Avatar          *string
	Background      *string
	BackgroundColor BackgroundColor
	PrimaryColor    PrimaryColor
	Email           string
	FullName        string
	InvitationCode  string
	Bio             *string
	Role            UserRole
	Credits         float64 `pg:",use_zero"`
	IsPrivate       bool    `pg:",use_zero"`
	IsVerified      bool    `pg:",use_zero"`
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	DeletedAt       *time.Time `pg:"-,soft_delete"`
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
