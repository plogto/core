package model

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	tableName struct{}   `sql:"user"`
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Fullname  *string    `json:"fullname"`
	Role      string     `json:"role"`
	IsPrivate bool       `json:"is_private"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}

func (u *User) GenToken() (*AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        u.ID,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "poster",
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
