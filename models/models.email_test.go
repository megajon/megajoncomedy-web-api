package models

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	emptyEmail := Email{
		ID:    1,
		Email: "",
	}

	invalidEmailFormat := Email{
		ID:    2,
		Email: "j",
	}

	validEmail := Email{
		ID:    3,
		Email: "jonathan.seubert@megajon.com",
	}

	longEmail := Email{
		ID:    4,
		Email: "pahoghaosgjdhgaohjlsdfhfgjfgjgfhsjfgjsgfghdfjfdgshfgjfgjfdhsfdhsfgjfsgjfjsgjfsgjssgsdhhdshfsdhsdfsdfhdshdfgdfgjfgjsfgjsfgjfgjfggjgfjdghgfhgdfhdfhdfhdfgsddhdhdfhdffhdfhdfdfgsdgdfhdsfgdsfsdfgdfgsdfgdfddhdfhdsgdfgdfgsdfgdfgdfgdfggsdfgsdfgdfgdgdsfgdfhdfhfhfghdfg@megajon.com",
	}

	err := ValidateEmail(emptyEmail)
	if err == nil {
		t.Error("email is blank and should not be allowed")
	}

	err = ValidateEmail(invalidEmailFormat)
	if err == nil {
		t.Error("invalid email format should not be allowed")
	}

	err = ValidateEmail(validEmail)
	if err != nil {
		t.Error("email is invalid")
	}

	err = ValidateEmail(longEmail)
	if err == nil {
		t.Error("email is too long and should not be allowed")
	}

}
