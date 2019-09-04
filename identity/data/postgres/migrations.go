package postgres

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

func Drop(db sqlx.Execer) error {
	_, err := db.Exec(`
		DROP TABLE IF EXISTS accounts
	`)
	if err != nil {
		log.Printf("Could not perform drops[%v]", err)
	}
	return err
}

func Migrate(db sqlx.Execer) error {
	migrations := []func(db sqlx.Execer) (sql.Result, error){
		mCreateAccounts,
	}
	for _, m := range migrations {
		result, err := m(db)
		if err != nil {
			return err
		}

		count, err := result.RowsAffected()
		if err != nil {
			log.Printf("Migrated, but couldn't get affected rows due to [%v]", err)
		}

		log.Printf("Migrated. %v rows affected", count)
	}
	return nil
}

func mCreateAccounts(db sqlx.Execer) (sql.Result, error) {
	return db.Exec(`
	CREATE TABLE IF NOT EXISTS accounts (
		account_id SERIAL PRIMARY KEY
		, username TEXT UNIQUE
		, password TEXT
	)
	`)
}
