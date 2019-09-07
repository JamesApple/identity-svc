package postgres_test

import (
	"log"
	"os"
	"root/data/postgres"
	"root/data/tests"
	"testing"

	"github.com/stretchr/testify/require"
)

var migrated = false

func TestRepo(t *testing.T) {
	ar, err := setup()
	require.NoError(t, err, "Could not setup test PG database")
	for _, test := range tests.AccountRepoTests {
		ar.DB.MustExec("TRUNCATE accounts RESTART IDENTITY CASCADE")
		test(t, ar)
	}
}

func setup() (*postgres.AccountRepo, error) {
	db, err := postgres.Connect("postgres://localhost/identity_test?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err := postgres.Migrate(db); err != nil {
		return nil, err
	}
	if !migrated {
		migrated = true
		postgres.Drop(db)
		postgres.Migrate(db)
	}
	ql := &postgres.InstrumentedDB{DB: db, Logger: log.New(os.Stdout, "TestDB", 0)}
	return &postgres.AccountRepo{InstrumentedDB: ql}, nil
}
