package postgres

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type InstrumentedDB struct {
	DB     *sqlx.DB
	Logger *log.Logger
}

func (q *InstrumentedDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	q.Logger.Print(query, args)
	return q.DB.Exec(query, args...)
}

func (q *InstrumentedDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	q.Logger.Print(query, args)
	return q.DB.Query(query, args...)
}

func (q *InstrumentedDB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	q.Logger.Print(query, args)
	return q.DB.Queryx(query, args...)
}

func (q *InstrumentedDB) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	q.Logger.Print(query, args)
	return q.DB.QueryRowx(query, args...)
}

func (q *InstrumentedDB) QueryRow(query string, args ...interface{}) *sql.Row {
	q.Logger.Print(query, args)
	return q.DB.QueryRow(query, args...)
}

func (q *InstrumentedDB) Stats() sql.DBStats {
	return q.DB.Stats()
}
