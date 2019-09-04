package mock

import (
	"database/sql"
	"fmt"
	"root"
)

func NewAccountRepo() *accountRepo {
	return &accountRepo{accounts: make([]*root.Account, 0)}
}

type accountRepo struct {
	accounts []*root.Account
}

func (a *accountRepo) Find(id int) (*root.Account, error) {
	for _, acc := range a.accounts {
		if acc.ID == id {
			return acc, nil
		}
	}
	return nil, nil
}

func (a *accountRepo) FindByUsername(username string) (*root.Account, error) {
	for _, acc := range a.accounts {
		if acc.Username == username {
			return acc, nil
		}
	}
	return nil, nil
}

func (a *accountRepo) Create(username string, password []byte) (*root.Account, error) {
	for _, acc := range a.accounts {
		if acc.Username == username {
			return nil, fmt.Errorf("Cannot use existing username %v", username)
		}
	}

	acc := &root.Account{
		ID:       len(a.accounts) + 1,
		Username: username,
		Password: password,
	}

	a.accounts = append(a.accounts, acc)
	return acc, nil
}

func (a *accountRepo) SetPassword(id int, password []byte) (bool, error) {
	acc, err := a.Find(id)
	if err != nil {
		return false, err
	}
	acc.Password = password
	return true, nil
}

func (a *accountRepo) SetUsername(id int, username string) (bool, error) {
	acc, err := a.Find(id)
	if err != nil {
		return false, err
	}
	acc.Username = username
	return true, nil
}

func (a *accountRepo) Stats() sql.DBStats {
	return sql.DBStats{OpenConnections: 1}
}

func duplicateAccount(acc *root.Account) *root.Account {
	return &root.Account{
		ID:       acc.ID,
		Password: acc.Password,
		Username: acc.Username,
	}
}
