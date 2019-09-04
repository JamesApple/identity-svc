package handlers

import (
	"net/http"
)

type PostAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func PostAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		WriteData(w, 200, "Hi, I am the handler")
	}
}
