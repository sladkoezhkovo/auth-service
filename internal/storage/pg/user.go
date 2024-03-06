package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

type userRepository struct {
	db *sqlx.DB
}

func (r *userRepository) Create(user *entity.User) error {
	if _, err := r.db.NamedExec("INSERT INTO users(email, password, role_id) VALUES (:email, :password, :role)", user); err != nil {
		return err
	}

	return nil
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

func NewUserRepository(conn *sqlx.DB) *userRepository {
	return &userRepository{
		db: conn,
	}
}
