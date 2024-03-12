package roleservice

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
	"github.com/sladkoezhkovo/auth-service/internal/service"
)

type RoleRepository interface {
	Create(role *entity.Role) error
	FindById(id int64) (*entity.Role, error)
	List(limit, offset int32) ([]*entity.Role, int, error)
	ListByName(name string, limit, offset int32) ([]*entity.Role, int, error)
	Update(role *entity.Role) error
	Delete(id int64) error
}

type roleService struct {
	repository RoleRepository
}

func New(r RoleRepository) *roleService {
	return &roleService{
		repository: r,
	}
}

func (s *roleService) Create(role *entity.Role) error {
	if err := s.repository.Create(role); err != nil {
		var pgerr *pq.Error
		if ok := errors.As(err, &pgerr); ok {
			switch pgerr.Code {
			case "23505":
				return service.ErrUniqueViolation
			}
			return pgerr
		}
		return err
	}

	return nil
}

func (s *roleService) FindById(id int64) (*entity.Role, error) {
	role, err := s.repository.FindById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}

	return role, nil
}

func (s *roleService) List(limit, offset int32) ([]*entity.Role, int, error) {
	entries, count, err := s.repository.List(limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, 0, service.ErrNotFound
		}
		return nil, 0, err
	}

	return entries, count, err
}

func (s *roleService) ListByName(name string, limit, offset int32) ([]*entity.Role, int, error) {
	entries, count, err := s.repository.ListByName(name, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, 0, service.ErrNotFound
		}
		return nil, 0, err
	}

	return entries, count, err
}

func (s *roleService) Update(role *entity.Role) error {
	if err := s.repository.Update(role); err != nil {
		return err
	}
	return nil
}

func (s *roleService) Delete(id int64) error {
	if err := s.repository.Delete(id); err != nil {
		return err
	}
	return nil
}
