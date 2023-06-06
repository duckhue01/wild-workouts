package cmerr

const (
	// common
	InternalServerError string = "internal-server-error"

	// token
	TokenIsExpired          string = "token-is-expired"
	MalformedToken          string = "malformed-jwt"
	EmptyBearerToken        string = "empty-bearer-token"
	InvalidSignature        string = "invalid-signature"
	UnexpectedSigningMethod string = "unexpected-signing-method"

	// auth middleware
	NoUserFound string = "no-user-found"
)
