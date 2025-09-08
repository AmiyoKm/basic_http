package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)
var DB_DSN = "postgres://basic_http:password@localhost/basic_http?sslmode=disable"

func initDB() *sql.DB {
	dsn := os.Getenv("DB_DSN")
	if dsn !=  "" {
		DB_DSN = dsn
	}
	db , err := sql.Open("postgres", DB_DSN )
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}