package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

type userRepository struct {
	connection *sqlx.DB
}

func (u *userRepository) Create(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Find(email string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Get(id int) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Update(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(conn *sqlx.DB) *userRepository {
	return &userRepository{
		connection: conn,
	}
}
