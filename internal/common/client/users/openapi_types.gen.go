// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package users

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// User defines model for User.
type User struct {
	Balance     int    `json:"balance"`
	DisplayName string `json:"displayName"`
	Role        string `json:"role"`
}
