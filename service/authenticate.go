package service

import (
	"root"
	"root/data"

	"golang.org/x/crypto/bcrypt"
)

func Authenticate(r data.AccountRepo, username string, password string) (*root.Account, error) {
	if username == "" || password == "" {
		return nil, FieldErrors{{Field: "base", Message: "Invalid Details"}}
	}

	account, err := r.FindByUsername(username)
	if err != nil {
		return nil, FieldErrors{{Field: "base", Message: "Invalid Details"}}
	}

	err = bcrypt.CompareHashAndPassword(account.Password, []byte(password))

	if err != nil {
		return nil, FieldErrors{{Field: "base", Message: "Invalid Details"}}
	}

	return account, nil
}
