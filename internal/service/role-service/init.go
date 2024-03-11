package roleservice

import (
	"errors"
	"github.com/lib/pq"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

type roleStorage interface {
	Create(role *entity.Role) error
	Find(name string) (*entity.Role, error)
	FindById(roleId int) (*entity.Role, error)
}

var (
	ErrUniqueViolation = errors.New("record already exists")
)

type roleService struct {
	storage roleStorage
}

func New(storage roleStorage) *roleService {
	return &roleService{
		storage: storage,
	}
}

func (r *roleService) Create(role *entity.Role) error {
	if err := r.storage.Create(role); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			switch pgerr.Code {
			case "23505":
				return ErrUniqueViolation
			}

			return pgerr
		}

		return err
	}

	return nil
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

func (r *roleService) FindById(roleId int) (*entity.Role, error) {
	role, err := r.storage.FindById(roleId)
	if err != nil {
		return nil, err
	}

	if role == nil {
		return role, errors.New("role not found")
	}

	return role, nil
}
