package util

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/graph/model"
)

func ParseJWTWithClaims(token string, claims *jwt.MapClaims) {
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func GenToken(userID pgtype.UUID) (*model.AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        convertor.UUIDToString(userID),
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
