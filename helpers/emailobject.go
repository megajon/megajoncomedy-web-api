package helpers

func CreateEmailObject(email string) (emailObject Email) {
	emailInput := Email{
		Email: email,
	}

	isEmailValid := ValidateEmail(emailInput)
	if isEmailValid != nil {
		emailInput.Email = "invalid"
	}

	return emailInput
}
