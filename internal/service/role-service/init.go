package roleservice

import (
	"errors"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

type roleStorage interface {
	Find(name string) (*entity.Role, error)
}

type roleService struct {
	storage roleStorage
}

func New(storage roleStorage) *roleService {
	return &roleService{
		storage: storage,
	}
}

func (r *roleService) Find(name string) (*entity.Role, error) {
	role, err := r.storage.Find(name)
	if err != nil {
		return nil, err
	}

	if role == nil {
		return role, errors.New("role not found")
	}

	return role, nil
}
