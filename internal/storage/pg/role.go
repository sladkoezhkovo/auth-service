package pg

import "database/sql"

type roleRepository struct {
	connection *sql.DB
}
