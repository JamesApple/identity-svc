package identity

import (
	"encoding/json"
	"net/http"
	"root/validator"
)

type httpServer struct {
	s *service
	v validator.Validator
}

func NewServer(s *service, v validator.Validator) *httpServer {
	return &httpServer{s: s, v: v}
}

func (h httpServer) Serve() {
	http.HandleFunc("/register", h.handleRegister)
	http.ListenAndServe(":8080", nil)
}

type registerRequest struct {
	Password string `json:"password" validate:"required,password"`
	Email    string `json:"email" validate:"required,email"`
}

type registerResponse struct {
	AccountID int    `json:"account_id"`
	Email     string `json:"email"`
}

func (h httpServer) handleRegister(w http.ResponseWriter, r *http.Request) {
	reg := &registerRequest{}

	err := json.NewDecoder(r.Body).Decode(reg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	errs, valid := h.v.ValidateS(reg)
	if valid != true {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}

	account, err := h.s.Register(reg.Email, reg.Password)
	if err != nil {
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte(err.Error()))
		return
	}

	resp := &registerResponse{Email: account.Email, AccountID: account.ID}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(500)
	}
}
