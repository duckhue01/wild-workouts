package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tribefintech/microservices/internal/common/cmerr"
	"github.com/tribefintech/microservices/internal/common/server/httperr"
)

type Parser interface {
	Parse(tokenString string) (*jwt.Token, error)
}

type User struct {
	UUID  string
	Email string
}

type AuthMiddleware struct {
	p  Parser
	wl []string
}

func NewAuthMiddleware(p Parser, whiteList []string) AuthMiddleware {
	return AuthMiddleware{
		p:  p,
		wl: whiteList,
	}
}

func (a AuthMiddleware) Middleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if contains(a.wl, r.URL.Path) {
			n.ServeHTTP(w, r)
			return
		}

		bearerToken := a.tokenFromRequest(r)

		if bearerToken == "" {
			httperr.Unauthorized(cmerr.EmptyBearerToken, nil, w, r)
			return
		}

		parsedTok, err := a.p.Parse(bearerToken)

		if parsedTok.Valid {
			if claims, ok := parsedTok.Claims.(jwt.MapClaims); ok {
				ctx = context.WithValue(ctx, userContextKey, User{
					UUID:  claims["sub"].(string),
					Email: claims["email"].(string),
				})
				r = r.WithContext(ctx)

				n.ServeHTTP(w, r)
				return
			}

			httperr.InternalError(cmerr.InternalServerError, err, w, r)
			return

		}

		if errors.Is(err, jwt.ErrTokenMalformed) {
			httperr.Unauthorized(cmerr.MalformedToken, err, w, r)
			return
		}

		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			httperr.Unauthorized(cmerr.InvalidSignature, err, w, r)
			return
		}

		if errors.Is(err, jwt.ErrTokenExpired) {
			httperr.Unauthorized(cmerr.TokenIsExpired, err, w, r)
			return
		}

		httperr.InternalError(cmerr.InternalServerError, err, w, r)
	})
}

func (c AuthMiddleware) tokenFromRequest(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}

	return r.URL.Query().Get("token")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
