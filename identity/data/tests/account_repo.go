package tests

import (
	"database/sql"
	"root/data"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type hasStats interface {
	Stats() sql.DBStats
}

func getOpenConnectionCount(repo data.AccountRepo) int {
	if st, ok := repo.(hasStats); ok {
		return st.Stats().OpenConnections
	}
	return 666
}

var AccountRepoTests = []func(*testing.T, data.AccountRepo){
	testCreate,
	testFind,
	testFindByUsername,
	testSetUsername,
	testSetPassword,
}

func testCreate(t *testing.T, repo data.AccountRepo) {
	account, err := repo.Create("test", []byte("password"))
	require.NoError(t, err)
	assert.Equal(t, 1, account.ID)
	assert.Equal(t, "test", account.Username)
	// Cannot duplicate name
	account, err = repo.Create("test", []byte("password"))
	require.Error(t, err)

	assert.Equal(t, 1, getOpenConnectionCount(repo))
}

func testFind(t *testing.T, repo data.AccountRepo) {
	_, err := repo.Create("test", []byte("password"))
	require.NoError(t, err)

	acc, err := repo.Find(1)
	require.NoError(t, err)
	require.NotNil(t, acc)
	assert.Equal(t, "test", acc.Username)

	acc, err = repo.Find(100)
	require.Nil(t, err)
	require.Nil(t, acc)

	assert.Equal(t, 1, getOpenConnectionCount(repo))
}

func testFindByUsername(t *testing.T, repo data.AccountRepo) {
	_, err := repo.Create("test", []byte("password"))
	require.NoError(t, err)

	acc, err := repo.FindByUsername("test")
	require.NoError(t, err)
	require.NotNil(t, acc)
	assert.Equal(t, "test", acc.Username)

	acc, err = repo.FindByUsername("Stephanie Worthington")
	require.Nil(t, err)
	require.Nil(t, acc)

	assert.Equal(t, 1, getOpenConnectionCount(repo))
}

func testSetUsername(t *testing.T, repo data.AccountRepo) {
	_, err := repo.Create("test", []byte("password"))
	require.NoError(t, err)

	valid, err := repo.SetUsername(1, "another")
	require.NoError(t, err)
	assert.Equal(t, valid, true)

	account, err := repo.FindByUsername("another")
	require.NoError(t, err)
	assert.Equal(t, account.Username, "another")

	assert.Equal(t, 1, getOpenConnectionCount(repo))
}

func testSetPassword(t *testing.T, repo data.AccountRepo) {
	_, err := repo.Create("test", []byte("password"))
	require.NoError(t, err)

	valid, err := repo.SetPassword(1, []byte("anotherPassword"))
	require.NoError(t, err)
	assert.Equal(t, valid, true)

	account, err := repo.Find(1)
	require.NoError(t, err)
	assert.Equal(t, account.Password, []byte("anotherPassword"))

	assert.Equal(t, 1, getOpenConnectionCount(repo))
}
