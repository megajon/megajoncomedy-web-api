package models

import "github.com/go-playground/validator/v10"

type Email struct {
	ID    int64  `bun:",pk,autoincrement"`
	Email string `json:"email" validate:"required,email,max=254"`
}

func ValidateEmail(email Email) (validated bool, err error) {
	v := validator.New()
	err = v.Struct(email)
	if err != nil {
		return false, err
	}
	return true, err
	// if err != nil {
	// 	for _, e := range err.(validator.ValidationErrors) {
	// 		fmt.Println(e)
	// 	}
	// }
}
