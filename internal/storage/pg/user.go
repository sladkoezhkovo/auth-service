package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) *userRepository {
	return &userRepository{
		db: conn,
	}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Get(
		user,
		`INSERT INTO users(email, password, role_id) VALUES ($1, $2, $3) RETURNING id, email, password, created_at, role_id as "role.id"`,
		user.Email,
		user.Password,
		user.Role.Id,
	)
}

func (r *userRepository) FindById(id int64) (*entity.User, error) {
	var user entity.User

	if err := r.db.Get(&user, `
SELECT 
	u.id, 
	u.email, 
	u.created_at, 
	r.id as "role.id", 
	r.name as "role.name"
FROM users u 
INNER JOIN public.roles r on r.id = u.role_id 
WHERE u.id = $1
LIMIT 1`,
		id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := r.db.Get(&user, `
SELECT 
	u.id, 
	u.email, 
	u.created_at, 
	r.id as "role.id", 
	r.name as "role.name"
FROM users u 
INNER JOIN public.roles r on r.id = u.role_id 
WHERE u.email = $1
LIMIT 1`,
		email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) List(limit, offset int32) ([]*entity.User, int, error) {
	var users []*entity.User
	if err := r.db.Select(&users, `
SELECT 
	u.id, u.email, u.created_at, r.name as "role.name"
FROM users u 
INNER JOIN public.roles r on r.id = u.role_id 
LIMIT $1 OFFSET $2`,
		limit, offset); err != nil {
		return nil, 0, err
	}
	return users, len(users), nil
}

func (r *userRepository) ListByRole(roleId int64, limit, offset int32) ([]*entity.User, int, error) {
	var users []*entity.User
	if err := r.db.Select(&users, `
SELECT 
	u.id, u.email, u.created_at, r.name as "role.name"
FROM users u 
INNER JOIN public.roles r on r.id = u.role_id 
WHERE u.role_id = $1
LIMIT $2 OFFSET $3`,
		roleId, limit, offset); err != nil {
		return nil, 0, err
	}
	return users, len(users), nil
}

func (r *userRepository) Update(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *userRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
