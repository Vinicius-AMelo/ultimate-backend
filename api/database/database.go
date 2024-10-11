package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() (*sql.DB, error) {
	var err error
	connString := "postgres://vini:vini@postgres:5432/vini?sslmode=disable"
	db, err = sql.Open("postgres", connString)

	return db, err
}

func GetDB() *sql.DB {
	return db
}
