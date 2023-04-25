package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/duckhue01/wild-workouts/internal/common/cmerr"
	"github.com/duckhue01/wild-workouts/internal/common/server/httperr"
	"github.com/golang-jwt/jwt"
)

type Parser interface {
	Parse(tokenString string) (*jwt.Token, error)
}

type User struct {
	UUID  string
	Email string
}

type AuthMiddleware struct {
	P Parser
}

func (a AuthMiddleware) Middleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		bearerToken := a.tokenFromHeader(r)
		if bearerToken == "" {
			httperr.Unauthorized(cmerr.EmptyBearerToken, nil, w, r)
			return
		}

		parsedTok, err := a.P.Parse(bearerToken)

		if parsedTok.Valid {
			if claims, ok := parsedTok.Claims.(jwt.MapClaims); ok {
				// it's always a good idea to use custom type as context value (in this case ctxKey)
				// because nobody from the outside of the package will be able to override/read this value
				ctx = context.WithValue(ctx, userContextKey, User{
					UUID:  claims["sub"].(string),
					Email: claims["email"].(string),
				})
				r = r.WithContext(ctx)

				n.ServeHTTP(w, r)
				return
			}

			httperr.InternalError(cmerr.InternalServerError, err, w, r)

		} else {
			if ve, ok := err.(*jwt.ValidationError); ok {

				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					httperr.Unauthorized(cmerr.MalformedToken, err, w, r)
					return
				}

				if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					httperr.Unauthorized(cmerr.TokenIsExpired, err, w, r)
					return
				}
			}

			httperr.InternalError(cmerr.InternalServerError, err, w, r)
			return

		}

	})
}

func (c AuthMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}

	return ""
}

type ctxKey int

const (
	userContextKey ctxKey = iota
)

var (
	// if we expect that the user of the function may be interested with concrete error,
	// it's a good idea to provide variable with this error
	NoUserInContextError = cmerr.NewAuthorizationError("no user in context", cmerr.NoUserFound)
)

func UserFromCtx(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userContextKey).(User)
	if ok {
		return u, nil
	}

	return User{}, NoUserInContextError
}
