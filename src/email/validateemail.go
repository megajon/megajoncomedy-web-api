package email

import (
	"github.com/go-playground/validator/v10"
	s "github.com/heroku/go-getting-started/src"
)

func ValidateEmail(email s.Email) (err error) {
	v := validator.New()
	err = v.Struct(email)
	return err
}
