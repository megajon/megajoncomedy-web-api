package helpers

import "github.com/go-playground/validator/v10"

func ValidateEmail(email Email) (err error) {
	v := validator.New()
	err = v.Struct(email)
	return err
}
