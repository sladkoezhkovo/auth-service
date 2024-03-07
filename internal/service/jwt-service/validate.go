package jwtservice

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
	"github.com/sladkoezhkovo/auth-service/internal/grpc/auth"
	"os"
	"strings"
)

func (s *jwtService) validate(tokenString string, secret string) (*entity.UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	}, jwt.WithExpirationRequired())
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, auth.ErrTokenExpired
		}
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
	claims, err := s.validate(tokenString, os.Getenv("JWT_REFRESH_SECRET"))
	if err != nil {
		return nil, err
	}

	validToken, err := s.storage.Get(claims.Email)
	if err != nil {
		return nil, err
	}

	if strings.Compare(validToken, tokenString) != 0 {
		return nil, fmt.Errorf("invalid refresh token")
	}

	return claims, nil
}

func (s *jwtService) ValidateAccess(tokenString string) (*entity.UserClaims, error) {
	return s.validate(tokenString, os.Getenv("JWT_ACCESS_SECRET"))
}
