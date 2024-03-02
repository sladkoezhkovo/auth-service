package jwtservice

import "github.com/golang-jwt/jwt/v5"

type userClaims struct {
	Role  int
	Email string
	jwt.RegisteredClaims
}
