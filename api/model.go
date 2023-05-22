package api

import (
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
