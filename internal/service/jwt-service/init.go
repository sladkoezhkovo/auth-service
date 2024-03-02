package jwtservice

import (
	"github.com/sladkoezhkovo/auth-service/internal/configs"
)

type TokenStorage interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, ttl int) error
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
	return s.storage.Set(email, token, 0)
}

func (s *jwtService) Clear(email string) error {
	return s.storage.Clear(email)
}
