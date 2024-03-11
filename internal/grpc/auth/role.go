package auth

import (
	"context"
	api "github.com/sladkoezhkovo/auth-service/api/auth"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

func (s *server) CreateRole(ctx context.Context, request *api.CreateRoleRequest) (*api.Empty, error) {

	role := &entity.Role{
		Name:      request.Name,
		Authority: request.Authority,
	}

	if err := s.roleService.Create(role); err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}
