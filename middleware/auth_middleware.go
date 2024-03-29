package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/database"
	"github.com/plogto/core/db"
	"github.com/plogto/core/validation"
)

type OnlineUserContext struct {
	User      db.User
	Token     string
	SocketID  string
	UserAgent string
}

var errAuthenticationFailed error = errors.New("Authentication failed")

func AuthMiddleware(users database.Users) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			user, err := users.GetUserByID(r.Context(), convertor.StringToUUID(claims["jti"].(string)))
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), constants.CURRENT_USER_KEY, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    StripBearerPrefixFromToken,
}

func StripBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv("JWT_SECRET"))
		return t, nil
	})

	return jwtToken, errors.Wrap(err, "parseToken error: ")
}

func GetCurrentUserFromCTX(ctx context.Context) (*db.User, error) {
	if ctx.Value(constants.CURRENT_USER_KEY) == nil {
		return nil, errAuthenticationFailed
	}

	user, ok := ctx.Value(constants.CURRENT_USER_KEY).(*db.User)

	if !ok || !validation.IsUserExists(user) {
		return nil, errAuthenticationFailed
	}

	return user, nil
}

func GetCurrentOnlineUserFromCTX(ctx context.Context) (*OnlineUserContext, error) {
	if ctx.Value(constants.CURRENT_ONLINE_USER_KEY) == nil {
		return nil, errAuthenticationFailed
	}

	onlineUser, ok := ctx.Value(constants.CURRENT_ONLINE_USER_KEY).(*OnlineUserContext)

	if !ok || onlineUser == nil {
		return nil, errAuthenticationFailed
	}

	return onlineUser, nil
}
