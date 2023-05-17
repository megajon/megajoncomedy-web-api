package email

import (
	"testing"
)

func TestCreateEmailObject(t *testing.T) {
	invalidEmail := "jonathan.seubert"
	validEmail := "jonathan.seubert@megajon.com"
	invalidEmailObject := CreateEmailObject(invalidEmail)
	validEmailObject := CreateEmailObject(validEmail)

	if invalidEmailObject.Email != "invalid" {
		t.Fatal("invalid email was allowed")
	}

	if validEmailObject.Email != validEmail {
		t.Fatal("invalid email detected.")
	}
}
