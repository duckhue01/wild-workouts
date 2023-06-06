package auth

import (
	"context"

	"github.com/tribefintech/microservices/internal/common/cmerr"
)

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
