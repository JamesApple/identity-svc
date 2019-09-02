package validator

import (
	"fmt"

	pkg "gopkg.in/go-playground/validator.v9"
)

type Validator interface {
	ValidateS(s interface{}) (map[string]string, bool)
}

type validator struct {
	*pkg.Validate
}

func New() *validator {
	val := &validator{Validate: pkg.New()}
	val.registerValidations()
	return val
}

func (v validator) registerValidations() {
	v.RegisterValidation("password", func(fl pkg.FieldLevel) bool {
		return len(fl.Field().String()) > 6
	})
}

func (v validator) ValidateS(s interface{}) (errs map[string]string, valid bool) {
	valid = true
	errs = make(map[string]string)

	err := v.Struct(s)

	if err != nil {
		valid = false
		for _, e := range err.(pkg.ValidationErrors) {
			errs[e.Namespace()] = fmt.Sprintf("invalid[%v]", e.Tag())
		}
	}
	return
}

var _ Validator = &validator{} // or &myType{} or [&]myType if scalar
