package postgres

import (
	"database/sql"
	"fmt"
	"root"

	"github.com/jmoiron/sqlx"
)

type AccountRepo struct {
	*QueryLogger
}

func (a *AccountRepo) Create(username string, password []byte) (*root.Account, error) {
	var id int
	account := &root.Account{
		Username: username,
		Password: password,
	}
	row := a.QueryRowx(`
		INSERT INTO accounts (
			username
			, password
		)
		VALUES ($1, $2)
		RETURNING account_id
	`, account.Username, account.Password)
	if row.Err() != nil {
		return nil, row.Err()
	}
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	account.ID = id
	return account, nil
}

func (a *AccountRepo) Find(id int) (*root.Account, error) {
	account := root.Account{}
	err := sqlx.Get(a, &account, "SELECT account_id, username, password FROM accounts WHERE account_id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *AccountRepo) FindByUsername(username string) (*root.Account, error) {
	account := root.Account{}
	err := sqlx.Get(a, &account, "SELECT account_id, username, password FROM accounts WHERE username = $1", username)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *AccountRepo) SetPassword(id int, password []byte) (bool, error) {
	return execOk(a.DB.Exec(`UPDATE accounts SET password = $1 WHERE account_id = $2`, password, id))
}

func (a *AccountRepo) SetUsername(id int, username string) (bool, error) {
	return execOk(a.DB.Exec(`UPDATE accounts SET username = $1 WHERE account_id = $2`, username, id))
}

func execOk(r sql.Result, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	count, err := r.RowsAffected()
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, fmt.Errorf("No records updated")
	}
	return true, nil
}
