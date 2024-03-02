package auth

import (
	"context"
	api "github.com/sladkoezhkovo/auth-service/api"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

type UserService interface {
	SignIn(email, password string) (*entity.User, error)
	SignUp(email, password string, role int) (*entity.User, error)
	Find(email string) (*entity.User, error)
}

type JwtService interface {
	Generate(u *entity.User) (*entity.Tokens, error)
	ValidateRefresh(token string) (*entity.UserClaims, error)
	ValidateAccess(token string) (*entity.UserClaims, error)
	Save(email, refresh string) error
	Clear(email string) error
}

type RoleService interface {
	Find(role string) (*entity.Role, error)
}

type server struct {
	userService UserService
	jwtService  JwtService
	roleService RoleService
	api.UnimplementedAuthServiceServer
}

func NewServer(user UserService, jwt JwtService, role RoleService) *server {
	return &server{
		userService: user,
		jwtService:  jwt,
		roleService: role,
	}
}

func (s *server) SignIn(ctx context.Context, request *api.SignInRequest) (*api.TokenResponse, error) {

	u, err := s.userService.SignIn(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	tokens, err := s.jwtService.Generate(u)
	if err != nil {
		return nil, err
	}

	if err := s.jwtService.Save(u.Email, tokens.Refresh); err != nil {
		return nil, err
	}

	return &api.TokenResponse{
		RefreshToken: tokens.Refresh,
		AccessToken:  tokens.Access,
	}, nil
}

func (s *server) SignUp(ctx context.Context, request *api.SignUpRequest) (*api.TokenResponse, error) {

	// TODO VALIDATE REQUEST

	role, err := s.roleService.Find(request.Role)
	if err != nil {
		return nil, err
	}

	u, err := s.userService.SignUp(request.Email, request.Password, role.Id)
	if err != nil {
		return nil, err
	}

	tokens, err := s.jwtService.Generate(u)
	if err != nil {
		return nil, err
	}

	if err := s.jwtService.Save(u.Email, tokens.Refresh); err != nil {
		return nil, err
	}

	return &api.TokenResponse{
		RefreshToken: tokens.Refresh,
		AccessToken:  tokens.Access,
	}, nil
}

func (s *server) Refresh(ctx context.Context, request *api.RefreshRequest) (*api.TokenResponse, error) {

	info, err := s.jwtService.ValidateRefresh(request.RefreshToken)
	if err != nil {
		return nil, err
	}

	u, err := s.userService.Find(info.Email)
	if err != nil {
		return nil, err
	}

	tokens, err := s.jwtService.Generate(u)
	if err != nil {
		return nil, err
	}

	if err := s.jwtService.Save(u.Email, tokens.Refresh); err != nil {
		return nil, err
	}

	return &api.TokenResponse{
		RefreshToken: tokens.Refresh,
		AccessToken:  tokens.Access,
	}, nil
}

func (s *server) Logout(ctx context.Context, request *api.LogoutRequest) (*api.Empty, error) {
	return &api.Empty{}, s.jwtService.Clear(request.Email)
}

func (s *server) Auth(ctx context.Context, request *api.AuthRequest) (*api.AuthResponse, error) {

	role, err := s.roleService.Find(request.Role)
	if err != nil {
		return nil, err
	}

	info, err := s.jwtService.ValidateAccess(request.AccessToken)
	if err != nil {
		// TODO add errors.Is() for expired token
		return nil, err
	}

	return &api.AuthResponse{
		Approved: role.Id == info.Role,
	}, nil
}
