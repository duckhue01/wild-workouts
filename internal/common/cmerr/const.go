package cmerr

const (
	// common
	InternalServerError string = "internal-server-error"

	// token
	TokenIsExpired   string = "token-is-expired"
	MalformedToken   string = "malformed-jwt"
	EmptyBearerToken string = "empty-bearer-token"

	// user
	NoUserFound string = "no-user-found"
)
