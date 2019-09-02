package root

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte("jwt_secret_key")

type Claims struct {
	AccountID int `json:"accountID"`
	jwt.StandardClaims
}

func NewToken(acc *Account) (Token, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		AccountID: acc.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return nil, err
	}

	return &tokenImpl{value: tokenString}, nil
}

type Token interface {
	String() string
}

type tokenImpl struct {
	value string
}

func (t tokenImpl) String() string {
	return t.value
}
