package service

import (
	"fmt"
	"strings"
)

var (
	ErrInvalid = "INVALID"
	ErrMissing = "MISSING"
)

type FieldErrors []FieldError
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

func (es FieldErrors) Error() string {
	errs := make([]string, len(es))
	for i, e := range es {
		errs[i] = e.Error()
	}

	return strings.Join(errs, ", ")
}
