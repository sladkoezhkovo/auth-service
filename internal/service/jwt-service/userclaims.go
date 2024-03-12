package jwtservice

import "github.com/golang-jwt/jwt/v5"

type userClaims struct {
	RoleId int64
	Email  string
	jwt.RegisteredClaims
}
