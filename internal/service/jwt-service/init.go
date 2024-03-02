package jwtservice

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sladkoezhkovo/auth-service/internal/configs"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
	"os"
	"time"
)

type userClaims struct {
	Role  int
	Email string
	jwt.RegisteredClaims
}

type TokenStorage interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Clear(key string) error
}

type jwtService struct {
	config  *configs.JwtConfig
	storage TokenStorage
}

func New(config *configs.JwtConfig, storage TokenStorage) *jwtService {
	return &jwtService{
		config:  config,
		storage: storage,
	}
}

func (s *jwtService) generateAccess(user *entity.User) (string, error) {
	claims := userClaims{
		Role:  user.Role,
		Email: user.Email,
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(s.config.AccessTTL)))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(os.Getenv("JWT_ACCESS_SECRET"))
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
	signed, err := token.SignedString(os.Getenv("JWT_ACCESS_SECRET"))
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

func (s *jwtService) validate(tokenString string, secret string) (*entity.UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(userClaims)
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

func (s *jwtService) Save(email, token string) error {
	return s.storage.Set(email, token)
}

func (s *jwtService) Clear(email string) error {
	return s.Clear(email)
}
