package entity

type Role struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}