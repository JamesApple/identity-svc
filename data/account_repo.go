package data

import "root"

type AccountRepo interface {
	Find(id int) (*root.Account, error)
	FindByUsername(username string) (*root.Account, error)

	Create(username string, password []byte) (*root.Account, error)

	SetPassword(id int, password []byte) (bool, error)
	SetUsername(id int, username string) (bool, error)
}
