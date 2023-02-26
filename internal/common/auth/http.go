package http

import (
	"net/http"
	"strings"
)

type UserCtx struct {
	UUID  string
	Email string
}

type CognitoHTTPMiddleware struct {
}

func (c CognitoHTTPMiddleware) Middleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()

		// token := c.tokenFromHeader(r)
		// if bearerToken == "" {
		// 	return
		// }

	})
}

func (c CognitoHTTPMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}

	return ""
}
