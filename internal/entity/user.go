package entity

import "time"

type User struct {
	Id        int64     `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	Role      `db:"role"`
}
