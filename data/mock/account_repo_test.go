package mock_test

import (
	"root/data/mock"
	"root/data/tests"
	"testing"
)

func TestMock(t *testing.T) {
	for _, test := range tests.AccountRepoTests {
		r := mock.NewAccountRepo()
		test(t, r)
	}
}
