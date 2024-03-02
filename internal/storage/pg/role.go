package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

type roleRepository struct {
	connection *sqlx.DB
}

func (r roleRepository) Find(name string) (*entity.Role, error) {
	//TODO implement me
	panic("implement me")
}

func NewRoleRepository(conn *sqlx.DB) *roleRepository {
	return &roleRepository{
		connection: conn,
	}
}
