package auth

import (
	"context"
	api "github.com/sladkoezhkovo/auth-service/api/auth"
	"github.com/sladkoezhkovo/auth-service/internal/converter"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

type RoleService interface {
	Create(role *entity.Role) error
	FindById(id int64) (*entity.Role, error)
	List(limit, offset int32) ([]*entity.Role, int, error)
	ListByName(name string, limit, offset int32) ([]*entity.Role, int, error)
	Update(role *entity.Role) error
	Delete(id int64) error
}

func (s *server) CreateRole(ctx context.Context, request *api.CreateRoleRequest) (*api.Role, error) {

	role := &entity.Role{
		Name:      request.Name,
		Authority: request.Authority,
	}

	if err := s.roleService.Create(role); err != nil {
		return nil, err
	}

	resp := converter.RoleFromEntityToDto(role)
	return resp, nil
}

func (s *server) FindByIdRole(ctx context.Context, request *api.FindRoleByIdRequest) (*api.Role, error) {
	role, err := s.roleService.FindById(request.Id)
	if err != nil {
		return nil, err
	}

	resp := converter.RoleFromEntityToDto(role)

	return resp, nil
}

func (s *server) ListRole(ctx context.Context, bounds *api.Bounds) (*api.ListRoleResponse, error) {
	roles, count, err := s.roleService.List(bounds.Limit, bounds.Offset)
	if err != nil {
		return nil, err
	}

	entries := make([]*api.Role, 0, count)

	for _, r := range roles {
		entries = append(entries, converter.RoleFromEntityToDto(r))
	}

	resp := &api.ListRoleResponse{
		Entries: entries,
		Count:   int64(count),
	}

	return resp, nil
}

func (s *server) ListRoleByName(ctx context.Context, request *api.ListRoleByNameRequest) (*api.ListRoleResponse, error) {
	roles, count, err := s.roleService.ListByName(request.Name, request.Bounds.Limit, request.Bounds.Offset)
	if err != nil {
		return nil, err
	}

	entries := make([]*api.Role, 0, count)

	for _, r := range roles {
		entries = append(entries, converter.RoleFromEntityToDto(r))
	}

	resp := &api.ListRoleResponse{
		Entries: entries,
		Count:   int64(count),
	}

	return resp, nil

}

func (s *server) UpdateRole(ctx context.Context, request *api.Role) (*api.Role, error) {
	role := converter.RoleFromDtoToEntity(request)
	if err := s.roleService.Update(role); err != nil {
		return nil, err
	}

	resp := converter.RoleFromEntityToDto(role)
	return resp, nil
}

func (s *server) DeleteRole(ctx context.Context, request *api.DeleteRoleRequest) (*api.Empty, error) {
	if err := s.roleService.Delete(request.Id); err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}
