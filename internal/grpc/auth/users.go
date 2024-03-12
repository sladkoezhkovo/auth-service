package auth

import (
	"context"
	api "github.com/sladkoezhkovo/auth-service/api/auth"
)

func (s *server) FindUserById(ctx context.Context, req *api.FindUserByIdRequest) (*api.UserDetails, error) {

	u, err := s.userService.FindById(req.Id)
	if err != nil {
		return nil, err
	}

	resp := &api.UserDetails{
		Id:        u.Id,
		Email:     u.Email,
		Role:      u.Role.Name,
		CreatedAt: u.CreatedAt.Unix(),
	}

	return resp, nil
}

func (s *server) ListUser(ctx context.Context, req *api.Bounds) (*api.ListUserResponse, error) {
	uu, count, err := s.userService.List(req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	resp := &api.ListUserResponse{
		Entries: make([]*api.User, 0, count),
		Count:   int32(count),
	}

	for _, u := range uu {
		resp.Entries = append(resp.Entries, &api.User{
			Id:    u.Id,
			Email: u.Email,
			Role:  u.Role.Name,
		})
	}

	return resp, nil
}

func (s *server) ListUserByRole(ctx context.Context, req *api.ListUserByRoleRequest) (*api.ListUserResponse, error) {
	uu, count, err := s.userService.ListByRole(req.RoleId, req.Bounds.Limit, req.Bounds.Offset)
	if err != nil {
		return nil, err
	}

	resp := &api.ListUserResponse{
		Entries: make([]*api.User, 0, count),
		Count:   int32(count),
	}

	for _, u := range uu {
		resp.Entries = append(resp.Entries, &api.User{
			Id:    u.Id,
			Email: u.Email,
			Role:  u.Role.Name,
		})
	}

	return resp, nil
}
