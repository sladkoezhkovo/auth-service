package entity

type Role struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Authority uint32 `db:"authority"`
}
