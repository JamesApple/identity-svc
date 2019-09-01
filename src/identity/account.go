package identity

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AccountRepo interface {
	Find(*Account) error
	FindByEmail(*Account) error

	EmailAvailable(*Account) (bool, error)

	Update(*Account) error
	Create(*Account) error
}

type Account struct {
	ID                int    `db:"account_id"`
	Email             string `db:"email"`
	EncryptedPassword string `db:"encrypted_password"`
}

func (a *Account) SetPassword(password string) error {
	encrypted_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	a.EncryptedPassword = string(encrypted_password)
	return nil
}

func (a Account) Authenticate(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(password)) == nil
}

func (a Account) String() string {
	return fmt.Sprintf("Account<ID: %v | Email: %v>", a.ID, a.Email)
}
