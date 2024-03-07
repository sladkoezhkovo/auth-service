package jwtservice

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
	"os"
)

func (s *jwtService) validate(tokenString string, secret string) (*entity.UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return nil, errors.New("unable to parse claims")
	}

	return &entity.UserClaims{
		Email: claims.Email,
		Role:  claims.Role,
	}, nil
}

func (s *jwtService) ValidateRefresh(tokenString string) (*entity.UserClaims, error) {
	return s.validate(tokenString, os.Getenv("JWT_REFRESH_SECRET"))
}

func (s *jwtService) ValidateAccess(tokenString string) (*entity.UserClaims, error) {
	return s.validate(tokenString, os.Getenv("JWT_ACCESS_SECRET"))
}
