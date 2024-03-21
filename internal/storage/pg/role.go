package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

type roleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(conn *sqlx.DB) *roleRepository {
	return &roleRepository{
		db: conn,
	}
}

func (r *roleRepository) Create(role *entity.Role) error {
	return r.db.Get(role,
		`INSERT INTO roles (name, authority) VALUES ($1, $2) RETURNING *`,
		role.Name,
		role.Authority,
	)
}

func (r *roleRepository) FindById(id int64) (*entity.Role, error) {
	var role entity.Role
	if err := r.db.Get(&role, "SELECT * FROM roles WHERE id=$1 LIMIT 1", id); err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) List(limit, offset int32) ([]*entity.Role, int, error) {
	var rr []*entity.Role
	if err := r.db.Select(&rr, "SELECT * FROM roles ORDER BY id LIMIT $1 OFFSET $2", limit, offset); err != nil {
		return nil, 0, err
	}

	var count int
	if err := r.db.Get(&count, "SELECT COUNT(id) FROM roles"); err != nil {
		return nil, 0, err
	}

	return rr, count, nil
}

func (r *roleRepository) ListByName(name string, limit, offset int32) ([]*entity.Role, int, error) {
	var rr []*entity.Role
	if err := r.db.Select(&rr, "SELECT * FROM roles WHERE name ILIKE $1 ORDER BY id LIMIT $2 OFFSET $3", "%"+name+"%", limit, offset); err != nil {
		return nil, 0, err
	}
	return rr, len(rr), nil
}

func (r *roleRepository) Update(role *entity.Role) error {
	return r.db.Get(
		role,
		`UPDATE roles SET name = $1, authority = $2 WHERE id = $3`,
		role.Name, role.Authority, role.Id,
	)
}

func (r *roleRepository) Delete(id int64) error {
	if _, err := r.db.Exec(`DELETE FROM roles WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}
