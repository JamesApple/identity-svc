package postgres

import (
	"root"

	"fmt"

	"github.com/jmoiron/sqlx"
)

type AccountRepo struct {
	*sqlx.DB
}

func (a *AccountRepo) Update(acc *root.Account) error {
	_, err := a.NamedExec(accountUpdate, acc)
	return err
}

func (a *AccountRepo) Create(acc *root.Account) error {
	statement, err := a.PrepareNamed(accountInsert)
	if err != nil {
		return err
	}

	_, err = statement.Exec(acc)

	return err
}

func (a *AccountRepo) Find(acc *root.Account) error {
	statement, err := a.PrepareNamed(accountFind)
	if err != nil {
		return err
	}
	if err := statement.QueryRow(acc).StructScan(acc); err != nil {
		return fmt.Errorf("Could not find account %v", acc)
	}

	return nil
}

func (a *AccountRepo) FindByEmail(acc *root.Account) error {
	statement, err := a.PrepareNamed(accountFindByEmail)
	if err != nil {
		return err
	}
	if err := statement.QueryRow(acc).StructScan(acc); err != nil {
		return fmt.Errorf("Could not find account %v", acc)
	}

	return nil
}

func (a *AccountRepo) EmailAvailable(acc *root.Account) (available bool, err error) {
	statement, err := a.PrepareNamed(accountEmailAvailable)
	if err != nil {
		return
	}

	err = statement.QueryRowx(acc).Scan(&available)
	return
}

func NewAccountRepo(db *sqlx.DB) *AccountRepo {
	return &AccountRepo{
		DB: db,
	}
}

const accountFind = `
SELECT
	account_id
	 , email
	, encrypted_password
FROM
	accounts
WHERE account_id = :account_id
`

const accountFindByEmail = `
SELECT
	account_id
	 , email
	, encrypted_password
FROM
	accounts
WHERE email = :email
`

const accountInsert = `
INSERT INTO 
accounts (
	email
	, encrypted_password
)
VALUES (
	:email
	, :encrypted_password
)
RETURNING account_id
`

const accountUpdate = `
UPDATE 
	accounts
SET 
	email = :email
	, encrypted_password = :encrypted_password
WHERE
	account_id = :account_id
`

const accountEmailAvailable = `
SELECT 
	COUNT(*) = 0 AS available
FROM 
	accounts
WHERE
	account_id != :account_id 
	AND email = :email
`
