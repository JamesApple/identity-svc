package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Open() *sqlx.DB {
	db, err := sqlx.Open("postgres", "postgres://localhost/identity?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
