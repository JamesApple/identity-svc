package handlers

import (
	"net/http"
	"root/cfg"
	"root/service"
)

type PostAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostAccountResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func PostAccount(app *cfg.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &PostAccountRequest{
			Username: "",
			Password: "",
		}
		err := ReadJSON(w, r, data)
		if err != nil {
			WriteErrors(w, 400, err)
			return
		}

		acc, err := service.CreateAccount(
			app.AccountRepo,
			data.Username,
			data.Password,
		)

		if err != nil {
			switch err.(type) {
			case service.FieldErrors:
				WriteErrors(w, 400, err)
			case service.FieldError:
				WriteErrors(w, 400, err)
			default:
				WriteErrors(w, 500, err)
			}
			return
		}

		WriteJSON(w, 200, &PostAccountResponse{
			ID:       acc.ID,
			Username: acc.Username,
		})
	}
}
