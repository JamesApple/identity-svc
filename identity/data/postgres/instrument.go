package postgres

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type QueryLogger struct {
	DB     *sqlx.DB
	Logger *log.Logger
}

func (q *QueryLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	q.Logger.Print(query, args)
	return q.DB.Exec(query, args...)
}

func (q *QueryLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	q.Logger.Print(query, args)
	return q.DB.Query(query, args...)
}

func (q *QueryLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	q.Logger.Print(query, args)
	return q.DB.Queryx(query, args...)
}

func (q *QueryLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	q.Logger.Print(query, args)
	return q.DB.QueryRowx(query, args...)
}

func (q *QueryLogger) QueryRow(query string, args ...interface{}) *sql.Row {
	q.Logger.Print(query, args)
	return q.DB.QueryRow(query, args...)
}

func (q *QueryLogger) Stats() sql.DBStats {
	return q.DB.Stats()
}
