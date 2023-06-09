package models

import "github.com/golang-jwt/jwt/v4"

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	UserId   int    `json:"userId"`
	jwt.RegisteredClaims
}
