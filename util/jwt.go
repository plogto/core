package util

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

func ParseJWTWithClaims(token string, claims *jwt.MapClaims) {
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
