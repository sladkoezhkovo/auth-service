package auth

import (
	"context"
	"errors"
	api "github.com/sladkoezhkovo/auth-service/api"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"os"
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
	Create(role *entity.Role) error
	Find(role string) (*entity.Role, error)
	FindById(roleId int) (*entity.Role, error)
}

type server struct {
	logger      *slog.Logger
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
		logger:      slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}
}

func (s *server) CreateRole(ctx context.Context, request *api.CreateRoleRequest) (*api.Empty, error) {

	l := s.logger.With(slog.String("handle", "CreateRole"))
	role := &entity.Role{
		Name:      request.Name,
		Authority: request.Authority,
	}

	if err := s.roleService.Create(role); err != nil {
		l.Error("error with creating role", slog.Any("err", err))
		return nil, err
	}

	l.Info("successfully created role", slog.String("role", role.Name))

	return &api.Empty{}, nil
}

func (s *server) SignIn(ctx context.Context, request *api.SignInRequest) (*api.TokenResponse, error) {

	l := s.logger.With(slog.String("handle", "SignIn"))
	u, err := s.userService.SignIn(request.Email, request.Password)
	if err != nil {
		l.Error("user cannot login", slog.String("email", request.Email), slog.Any("err", err))
		if errors.Is(err, ErrInvalidPassword) {
			return nil, status.Errorf(codes.Canceled, "invalid email or password")
		}
		return nil, err
	}

	tokens, err := s.jwtService.Generate(u)
	if err != nil {
		l.Error("tokens cant be generated", slog.Any("err", err))
		return nil, err
	}

	if err := s.jwtService.Save(u.Email, tokens.Refresh); err != nil {
		l.Error("refresh token unable to save", slog.String("email", u.Email), slog.String("refresh-token", tokens.Refresh), slog.Any("err", err))
		return nil, err
	}

	l.Info("successfully signin", slog.Any("user", u))

	return &api.TokenResponse{
		RefreshToken: tokens.Refresh,
		AccessToken:  tokens.Access,
	}, nil
}

func (s *server) SignUp(ctx context.Context, request *api.SignUpRequest) (*api.TokenResponse, error) {

	// TODO VALIDATE REQUEST

	l := s.logger.With(slog.String("handle", "SignIn"))
	role, err := s.roleService.Find(request.Role)
	if err != nil {
		l.Error("role unable to found", slog.String("role", request.Role), slog.Any("err", err))
		return nil, err
	}
	l.Debug("role found", slog.Any("role", role))

	u, err := s.userService.SignUp(request.Email, request.Password, role.Id)
	if err != nil {
		l.Error("unable to signup", slog.Any("req", request), slog.Int("role_id", role.Id), slog.Any("err", err))
		return nil, err
	}
	l.Debug("user signed up", slog.Any("user", u))

	tokens, err := s.jwtService.Generate(u)
	if err != nil {
		l.Error("tokens cant be generated", slog.Any("err", err))
		return nil, err
	}
	l.Error("tokens generated", slog.Any("tokens", tokens))

	if err := s.jwtService.Save(u.Email, tokens.Refresh); err != nil {
		l.Error("refresh token unable to save", slog.String("email", u.Email), slog.String("refresh-token", tokens.Refresh), slog.Any("err", err))
		return nil, err
	}

	l.Info("successfully signin", slog.Any("tokens", tokens))

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

	userRole, err := s.roleService.FindById(info.Role)
	if err != nil {
		// TODO add errors.Is() for expired token
		return nil, err
	}

	return &api.AuthResponse{
		Approved: role.Authority >= userRole.Authority,
	}, nil
}
