package entity

type Role struct {
	Id        int64  `db:"id"`
	Name      string `db:"name"`
	Authority int32  `db:"authority"`
}
