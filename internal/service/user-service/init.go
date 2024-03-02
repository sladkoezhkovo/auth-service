package userservice

import (
	"github.com/sladkoezhkovo/auth-service/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRepository interface {
	Create(user *entity.User) error
	Find(email string) (*entity.User, error)
	Get(id int) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id int) error
}

type RoleFinder interface {
	FindId(role string) (int, error)
}

type userService struct {
	repository UserRepository
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
		Email:     email,
		Password:  string(hash),
		Role:      role,
		CreatedAt: time.Now(),
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

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
