package jwtservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
	"os"
	"time"
)

func (s *jwtService) generateAccess(user *entity.User) (string, error) {
	claims := userClaims{
		Role:  user.Role,
		Email: user.Email,
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(s.config.AccessTTL)))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (s *jwtService) generateRefresh(user *entity.User) (string, error) {
	claims := userClaims{
		Role:  user.Role,
		Email: user.Email,
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(s.config.RefreshTTL)))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (s *jwtService) Generate(user *entity.User) (*entity.Tokens, error) {

	refresh, err := s.generateRefresh(user)
	if err != nil {
		return nil, err
	}

	access, err := s.generateAccess(user)
	if err != nil {
		return nil, err
	}

	return &entity.Tokens{
		Access:  access,
		Refresh: refresh,
	}, nil
}
