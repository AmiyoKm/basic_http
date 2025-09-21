package infra

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AmiyoKm/basic_http/config"
	_ "github.com/lib/pq"
)

func InitDB(cfg config.DbConfig) (*sql.DB, error) {
	DB_DSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
