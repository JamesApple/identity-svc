package data

import "github.com/lib/pq"

func IsUniquenessError(e error) bool {
	switch e.(type) {
	case *pq.Error:
		return true
	default:
		return false
	}
}
