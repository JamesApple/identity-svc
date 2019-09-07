package postgres

import (
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(str string) (*sqlx.DB, error) {
	dbURL, err := url.Parse(str)
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Open("postgres", dbURL.String())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
