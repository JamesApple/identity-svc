package root

type Account struct {
	ID       int `db:"account_id"`
	Username string
	Password []byte
}
