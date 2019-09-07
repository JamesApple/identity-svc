package handlers

import (
	"encoding/json"
	"net/http"
	"root/service"
)

type ResponseShape struct {
	Result interface{} `json:"result"`
}

type ErrorShape struct {
	Errors interface{} `json:"errors"`
}

func WriteData(w http.ResponseWriter, code int, data interface{}) {
	WriteJSON(w, code, &ResponseShape{Result: data})
}

func WriteErrors(w http.ResponseWriter, code int, data interface{}) {
	WriteJSON(w, code, &ErrorShape{Errors: data})
}

func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(j)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, readable interface{}) error {
	err := json.NewDecoder(r.Body).Decode(readable)
	if err != nil {
		return service.FieldErrors{
			service.FieldError{
				Field:   "Base",
				Message: "Invalid JSON",
			},
		}
	}
	return nil
}
