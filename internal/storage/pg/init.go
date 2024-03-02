package pg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sladkoezhkovo/auth-service/internal/configs"
	"os"
)

func Setup(config *configs.SqlConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), config.Db, config.TLS))
	if err != nil {
		return nil, err
	}

	return db, nil
}
