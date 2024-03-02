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

func (u *userService) SignUp(email, password, role string) (*entity.User, error) {
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
}
