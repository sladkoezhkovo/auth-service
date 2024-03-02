package jwtservice

import "github.com/sladkoezhkovo/auth-service/internal/configs"

type jwtService struct {
	config *configs.JwtConfig
}

func New(config *configs.JwtConfig) *jwtService {
	return &jwtService{
		config: config,
	}
}

func (s *jwtService) GenerateToken(payload interface{}) {

}
