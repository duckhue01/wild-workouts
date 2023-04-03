// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package ports

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Demo defines model for Demo.
type Demo struct {
	Name string `json:"name"`
}
