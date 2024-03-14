package userservice

import (
	"database/sql"
	"errors"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
	"github.com/sladkoezhkovo/auth-service/internal/grpc/auth"
	"github.com/sladkoezhkovo/auth-service/internal/service"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindById(id int64) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	List(limit, offset int32) ([]*entity.User, int, error)
	ListByRole(roleId int64, limit, offset int32) ([]*entity.User, int, error)
	Update(user *entity.User) error
	Delete(id int) error
}

type userService struct {
	repository UserRepository
}

func New(repo UserRepository) *userService {
	return &userService{
		repository: repo,
	}
}

func (s *userService) FindByEmail(email string) (*entity.User, error) {
	u, err := s.repository.FindByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}

	return u, nil
}

func (s *userService) FindById(id int64) (*entity.User, error) {
	u, err := s.repository.FindById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}

	return u, nil
}

func (s *userService) List(limit, offset int32) ([]*entity.User, int, error) {
	uu, count, err := s.repository.List(limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, 0, service.ErrNotFound
		}
		return nil, 0, err
	}

	return uu, count, nil
}

func (s *userService) ListByRole(roleId int64, limit, offset int32) ([]*entity.User, int, error) {
	uu, count, err := s.repository.ListByRole(roleId, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, 0, service.ErrNotFound
		}
		return nil, 0, err
	}

	return uu, count, nil
}

func (s *userService) SignUp(user *entity.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	if err := s.repository.Create(user); err != nil {
		return err
	}

	return nil
}

func (s *userService) SignIn(user *entity.User) (*entity.User, error) {
	candidate, err := s.repository.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(candidate.Password), []byte(user.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, auth.ErrInvalidPassword
		}
		return nil, err
	}

	return candidate, nil
}
