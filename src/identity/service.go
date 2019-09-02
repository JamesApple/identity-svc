package identity

import (
	"root"
)

func NewService(repo root.AccountRepo) *service {
	return &service{AccountRepo: repo}
}

type service struct {
	AccountRepo root.AccountRepo
}

func (i *service) Register(email string, password string) (*root.Account, error) {
	account := &root.Account{Email: email}

	available, err := i.AccountRepo.EmailAvailable(account)
	if err != nil {
		return nil, err
	}

	if available == false {
		return nil, root.Err(root.ErrNotAvailable, "Email already taken")
	}

	err = account.SetPassword(password)
	if err != nil {
		return nil, err
	}

	err = i.AccountRepo.Create(account)
	if err != nil {
		return nil, err
	}
	err = i.AccountRepo.FindByEmail(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (i *service) Login(email string, password string) (root.Token, error) {
	account := &root.Account{Email: email}

	err := i.AccountRepo.FindByEmail(account)
	if err != nil {
		return nil, err
	}

	if account.Authenticate(password) == false {
		return nil, root.Err(root.ErrNotFound, "Could not authenticate")
	}

	return root.NewToken(account)
}

func (i *service) Refresh(token string) (root.Token, error) {
	panic("not implemented")
}
