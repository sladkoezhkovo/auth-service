package entity

import "time"

type User struct {
	Id        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      int       `db:"role_id"`
	CreatedAt time.Time `db:"created_at"`
}
