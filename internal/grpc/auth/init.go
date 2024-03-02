package auth

import (
	"context"
	api "github.com/sladkoezhkovo/auth-service/api/authservice"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

type UserService interface {
	SignIn(email, password string) error
	SignUp(email, password string, role int) (*entity.User, error)
	Find(email string) (*entity.User, error)
}

type JwtService interface {
	GenerateAccess(payload interface{}) (string, error)
	GenerateRefresh(payload interface{}) (string, error)
	Validate(token string) (string, error)
	Save(refresh string) error
	Clear(email string) error
}

type RoleService interface {
	GetId(role string) (int, error)
}

type server struct {
	userService UserService
	jwtService  JwtService
	roleService RoleService
}

func NewServer(user UserService, jwt JwtService, role RoleService) *server {
	return &server{
		userService: user,
		jwtService:  jwt,
		roleService: role,
	}
}

func (s *server) SignIn(ctx context.Context, request *api.SignInRequest) (*api.TokenResponse, error) {

	// TODO VALIDATE REQUEST

	//return response, nil
}

func (s *server) SignUp(ctx context.Context, request *api.SignUpRequest) (*api.TokenResponse, error) {

	// TODO VALIDATE REQUEST

	roleId, err := s.roleService.GetId(request.Role)
	if err != nil {
		return nil, err
	}

	u, err := s.userService.SignUp(request.Email, request.Password, roleId)
	if err != nil {
		return nil, err
	}

	access, err := s.jwtService.GenerateAccess(u)
	if err != nil {
		return nil, err
	}

	refresh, err := s.jwtService.GenerateRefresh(u)
	if err != nil {
		return nil, err
	}

	if err := s.jwtService.Save(refresh); err != nil {
		return nil, err
	}

	response := &api.TokenResponse{
		RefreshToken: refresh,
		AccessToken:  access,
	}

	return response, nil
}

func (s *server) Refresh(ctx context.Context, request *api.RefreshRequest) (*api.TokenResponse, error) {

	email, err := s.jwtService.Validate(request.RefreshToken)
	if err != nil {
		return nil, err
	}

	u, err := s.userService.Find(email)
	if err != nil {
		return nil, err
	}

	refresh, err := s.jwtService.GenerateRefresh(u)
	if err != nil {
		return nil, err
	}

	access, err := s.jwtService.GenerateAccess(u)
	if err != nil {
		return nil, err
	}

	if err := s.jwtService.Save(refresh); err != nil {
		return nil, err
	}

	return &api.TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *server) Logout(ctx context.Context, request *api.LogoutRequest) (*api.Empty, error) {
	return &api.Empty{}, s.jwtService.Clear(request.Email)
}
