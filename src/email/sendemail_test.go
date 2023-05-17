package email

import (
	"testing"
)

func Test_SendEmail(t *testing.T) {
	sender := "noreply@megajon.com"
	newSubscriberEmail := "test.email@megajon.com"
	testEmail := SendNewSubscriberEmail(newSubscriberEmail)
	if testEmail.Sender != sender {
		t.Error("Sender does not match")
	}
}
