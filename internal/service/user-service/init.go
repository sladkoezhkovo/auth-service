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

func (s *userService) SignIn(user *entity.User) error {
	candidate, err := s.repository.Find(user.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(candidate.Password), []byte(user.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return auth.ErrInvalidPassword
		}
		return err
	}

	return nil
}
