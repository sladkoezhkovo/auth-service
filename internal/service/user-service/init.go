package userservice

import (
	"errors"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
	"github.com/sladkoezhkovo/auth-service/internal/grpc/auth"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user *entity.User) error
	Find(email string) (*entity.User, error)
	Get(id int) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id int) error
}

type userService struct {
	repository UserRepository
}

func (s *userService) Find(email string) (*entity.User, error) {

	u, err := s.repository.Find(email)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("user not found")
	}

	return u, nil
}

func New(repo UserRepository) *userService {
	return &userService{
		repository: repo,
	}
}

func (s *userService) SignUp(email, password string, role int) (*entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Email:    email,
		Password: string(hash),
		Role:     role,
	}

	if err := s.repository.Create(user); err != nil {
		return nil, err
	}

	return user, err
}

func (s *userService) SignIn(email, password string) (*entity.User, error) {
	user, err := s.repository.Find(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, auth.ErrInvalidPassword
		}
		return nil, err
	}

	return user, nil
}
