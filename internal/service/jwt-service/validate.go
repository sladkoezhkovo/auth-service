package jwtservice

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
	"os"
	"time"
)

func (s *jwtService) validate(tokenString string, secret string) (*entity.UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	}, jwt.WithExpirationRequired())
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("unathorized")
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return nil, errors.New("unable to parse claims")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token expired")
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
