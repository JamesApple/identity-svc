package service

import (
	"root"
	"root/data"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(repo data.AccountRepo, username string, password string) (*root.Account, error) {
	username = strings.TrimSpace(username)

	errs := FieldErrors{}
	fieldErr := ValidateUsername(username)
	if fieldErr != nil {
		errs = append(errs, *fieldErr)
	}

	fieldErr = ValidatePassword(password)
	if fieldErr != nil {
		errs = append(errs, *fieldErr)
	}

	if len(errs) > 0 {
		return nil, errs
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	acc, err := repo.Create(username, hash)

	if err != nil {
		if data.IsUniquenessError(err) {
			return nil, FieldErrors{
				FieldError{
					Field:   "username",
					Message: "TAKEN",
				},
			}
		}
		return nil, err
	}

	return acc, nil
}
