package auth

import "errors"

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrTokenExpired    = errors.New("token expired")
)
