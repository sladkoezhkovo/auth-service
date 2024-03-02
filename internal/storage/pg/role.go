package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

type roleRepository struct {
	db *sqlx.DB
}

func (r *roleRepository) Create(role *entity.Role) error {
	return r.db.Get(&role.Id, "INSERT INTO role(name) VALUES ($1)", role.Name)
}

func (r *roleRepository) Find(name string) (*entity.Role, error) {

	var role []entity.Role

	if err := r.db.Select(&role, "SELECT * FROM role WHERE name=$1", name); err != nil {
		return nil, err
	}

	if len(role) < 1 {
		//return nil, errors.New("role not found")
		return nil, nil
	}

	return &role[0], nil
}

func NewRoleRepository(conn *sqlx.DB) *roleRepository {
	return &roleRepository{
		db: conn,
	}
}
