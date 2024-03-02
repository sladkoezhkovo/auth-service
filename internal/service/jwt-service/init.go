package jwtservice

import (
	"github.com/sladkoezhkovo/auth-service/internal/configs"
)

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

func (s *jwtService) Save(email, token string) error {
	return s.storage.Set(email, token)
}

func (s *jwtService) Clear(email string) error {
	return s.storage.Clear(email)
}
