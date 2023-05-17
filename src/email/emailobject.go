package email

import s "github.com/heroku/go-getting-started/src"

func CreateEmailObject(email string) (emailObject s.Email) {
	emailInput := s.Email{
		Email: email,
	}

	isEmailValid := ValidateEmail(emailInput)
	if isEmailValid != nil {
		emailInput.Email = "invalid"
	}

	return emailInput
}
