package util

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/plogto/core/graph/model"
)

func ParseJWTWithClaims(token string, claims *jwt.MapClaims) {
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func GenToken(userID uuid.UUID) (*model.AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        userID.String(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "plog",
	})

	token, err := t.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &model.AuthToken{
		Token:     token,
		ExpiredAt: expiredAt,
	}, nil
}
