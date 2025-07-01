package helper

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func CustomValidation() *validator.Validate {
	v := validator.New()

	_ = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		var hasUpper, hasLower, hasNumber bool

		for _, c := range password {
			switch {
			case unicode.IsUpper(c):
				hasUpper = true
			case unicode.IsLower(c):
				hasLower = true
			case unicode.IsDigit(c):
				hasNumber = true
			}
		}

		return len(password) >= 6 && hasUpper && hasLower && hasNumber
	})

	return v
}
