package entity

import "time"

type User struct {
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      int       `db:"role"`
	CreatedAt time.Time `db:"created_at"`
}
