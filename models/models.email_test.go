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

	testEmptyEmail, err := ValidateEmail(emptyEmail)
	if testEmptyEmail == true && err == nil {
		t.Fatal("email is blank and should not be allowed")
	}

	testInvalidEmailFormat, err := ValidateEmail(invalidEmailFormat)
	if testInvalidEmailFormat == true && err == nil {
		t.Fatal("invalid email format should not be allowed")
	}

	testValidEmail, err := ValidateEmail(validEmail)
	if testValidEmail != true && err != nil {
		t.Fatal("email is invalid")
	}

	testLongEmail, err := ValidateEmail(longEmail)
	if testLongEmail == true && err == nil {
		t.Fatal("email is too long and should not be allowed")
	}

}
