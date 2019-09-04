package handlers

import (
	"encoding/json"
	"net/http"
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
