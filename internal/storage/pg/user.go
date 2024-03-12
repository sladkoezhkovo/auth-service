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

func (r *userRepository) Find(email string) (*entity.User, error) {
	var users []entity.User

	if err := r.db.Select(&users, "SELECT * FROM users WHERE email=$1", email); err != nil {
		return nil, err
	}

	if len(users) < 1 {
		//return nil, errors.New("role not found")
		return nil, nil
	}

	return &users[0], nil
}

func (r *userRepository) Get(id int) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepository) Update(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *userRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
