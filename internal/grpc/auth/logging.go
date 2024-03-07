package auth

import (
	"context"
	"github.com/sladkoezhkovo/auth-service/api"
	"google.golang.org/grpc/status"
	"log/slog"
	"os"
)

type loggingServer struct {
	logger *slog.Logger
	srv    *server
	api.UnimplementedAuthServiceServer
}

func NewLoggingServer(user UserService, jwt JwtService, role RoleService) *loggingServer {
	return &loggingServer{
		srv:    NewServer(user, jwt, role),
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}
}

func (s *loggingServer) log(handle string, err error) *slog.Logger {
	l := s.logger.With(slog.String("handle", handle))
	if err != nil {
		if e, ok := status.FromError(err); ok {
			l = l.With(
				slog.Int("code", int(e.Code())),
				slog.String("status", e.Code().String()),
			)
			l.Error(e.Message())
		}
	} else {
		l.Info("OK")
	}
	return l
}

func (s *loggingServer) CreateRole(ctx context.Context, request *api.CreateRoleRequest) (*api.Empty, error) {
	r, err := s.srv.CreateRole(ctx, request)
	s.log("CreateRole", err)
	return r, err
}
func (s *loggingServer) SignIn(ctx context.Context, request *api.SignInRequest) (*api.TokenResponse, error) {
	r, err := s.srv.SignIn(ctx, request)
	s.log("SignIn", err)
	return r, err
}
func (s *loggingServer) SignUp(ctx context.Context, request *api.SignUpRequest) (*api.TokenResponse, error) {
	r, err := s.srv.SignUp(ctx, request)
	s.log("SignUp", err)
	return r, err
}
func (s *loggingServer) Refresh(ctx context.Context, request *api.RefreshRequest) (*api.TokenResponse, error) {
	r, err := s.srv.Refresh(ctx, request)
	s.log("Refresh", err)
	return r, err
}
func (s *loggingServer) Logout(ctx context.Context, request *api.LogoutRequest) (*api.Empty, error) {
	r, err := s.srv.Logout(ctx, request)
	s.log("Logout", err)
	return r, err
}
func (s *loggingServer) Auth(ctx context.Context, request *api.AuthRequest) (*api.AuthResponse, error) {
	r, err := s.srv.Auth(ctx, request)
	s.log("Auth", err)
	return r, err
}
