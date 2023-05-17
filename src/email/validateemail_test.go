package email

import (
	"testing"

	s "github.com/heroku/go-getting-started/src"
)

func TestValidateEmail(t *testing.T) {
	emptyEmail := s.Email{
		ID:    1,
		Email: "",
	}

	invalidEmailFormat := s.Email{
		ID:    2,
		Email: "j",
	}

	validEmail := s.Email{
		ID:    3,
		Email: "jonathan.seubert@megajon.com",
	}

	longEmail := s.Email{
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
