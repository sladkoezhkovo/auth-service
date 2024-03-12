package auth

import (
	"context"
	"errors"
	api "github.com/sladkoezhkovo/auth-service/api/auth"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	SignIn(user *entity.User) error
	SignUp(user *entity.User) error
	FindById(id int64) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	List(limit, offset int32) ([]*entity.User, int, error)
	ListByRole(roleId int64, limit, offset int32) ([]*entity.User, int, error)
}

type JwtService interface {
	Generate(u *entity.User) (*entity.Tokens, error)
	ValidateRefresh(token string) (*entity.UserClaims, error)
	ValidateAccess(token string) (*entity.UserClaims, error)
	Save(email, refresh string) error
	Clear(email string) error
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

	user := &entity.User{
		Email:    request.Email,
		Password: request.Password,
		Role:     entity.Role{},
	}

	if err := s.userService.SignIn(user); err != nil {
		if errors.Is(err, ErrInvalidPassword) {
			return nil, status.Errorf(codes.Canceled, "invalid email or password")
		}
		return nil, err
	}

	tokens, err := s.jwtService.Generate(user)
	if err != nil {
		return nil, err
	}

	if err := s.jwtService.Save(user.Email, tokens.Refresh); err != nil {
		return nil, err
	}

	return &api.TokenResponse{
		RefreshToken: tokens.Refresh,
		AccessToken:  tokens.Access,
	}, nil
}

func (s *server) SignUp(ctx context.Context, request *api.SignUpRequest) (*api.TokenResponse, error) {

	user := &entity.User{
		Email:    request.Email,
		Password: request.Password,
		Role: entity.Role{
			Id: request.RoleId,
		},
	}

	if err := s.userService.SignUp(user); err != nil {
		return nil, err
	}

	tokens, err := s.jwtService.Generate(user)
	if err != nil {
		return nil, err
	}

	if err := s.jwtService.Save(user.Email, tokens.Refresh); err != nil {
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
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	u, err := s.userService.FindByEmail(info.Email)
	if err != nil {
		return nil, status.Errorf(codes.Canceled, err.Error())
	}

	tokens, err := s.jwtService.Generate(u)
	if err != nil {
		return nil, status.Errorf(codes.Canceled, err.Error())
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
	token, err := s.jwtService.ValidateAccess(request.AccessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	return &api.Empty{}, s.jwtService.Clear(token.Email)
}

func (s *server) Auth(ctx context.Context, request *api.AuthRequest) (*api.AuthResponse, error) {

	role, err := s.roleService.FindById(request.RoleId)
	if err != nil {
		return nil, err
	}

	info, err := s.jwtService.ValidateAccess(request.AccessToken)
	if err != nil {
		if errors.Is(err, ErrTokenExpired) {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}
		return nil, status.Errorf(codes.Canceled, err.Error())
	}

	userRole, err := s.roleService.FindById(info.Role)
	if err != nil {
		// TODO add errors.Is() for expired token
		return nil, status.Errorf(codes.Canceled, err.Error())
	}

	response := &api.AuthResponse{}

	if role.Authority == userRole.Authority {
		response.Approved = role.Id == userRole.Id
	} else {
		response.Approved = role.Authority > userRole.Authority
	}

	return response, nil
}
