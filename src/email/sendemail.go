package email

import (
	"fmt"

	s "github.com/heroku/go-getting-started/src"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	Sender  = "megajon@megajon.com"
	CharSet = "UTF-8"
)

func SendNewSubscriberEmail(newSubscriberEmail string) s.OutgoingEmail {
	subject := "Megajon has a new subscriber."
	htmlBody := fmt.Sprintf("%s has joined the ranks of Megajon!", newSubscriberEmail)
	textBody := "This text body is from megajon.com"

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(Sender),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(htmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(textBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(Sender),
	}

	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	fmt.Println("Email Sent to address: " + Sender)
	fmt.Println(result)
	emailObject := s.OutgoingEmail{
		Sender:    Sender,
		Recipient: Sender,
		HtmlBody:  htmlBody,
		TextBody:  textBody,
		CharSet:   CharSet,
	}
	return emailObject

}

func SendWelcomeEmail(newSubscriberEmail string) s.OutgoingEmail {
	subject := "Welcome to the world of Megajon!"
	htmlBody := `<h1>Thank you for joining the ranks of Megajon! I hope to keep you entertained
	as you follow me through my journey through the comedy world and beyond. You'll
	receive an email whenever I post new content.</h1>
	<h1>If you wish you can remove yourself from the mailing list anytime by using
	this <a href='https://megajon.com/unsubscribe'>unsubscribe link</a>. You'll be making a lame choice and will miss all the fun,
	but Megajon doesn't hold hostages.</h1>`
	textBody := ""

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(newSubscriberEmail),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(htmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(textBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(Sender),
	}

	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	fmt.Println("Email Sent to address: " + newSubscriberEmail)
	fmt.Println(result)
	emailObject := s.OutgoingEmail{
		Sender:    Sender,
		Recipient: newSubscriberEmail,
		HtmlBody:  htmlBody,
		TextBody:  textBody,
		CharSet:   CharSet,
	}
	return emailObject

}

func UnsubscribeEmail(unsubscriberEmail string) s.OutgoingEmail {
	subject := fmt.Sprintf("%s has unsubscribed.", unsubscriberEmail)
	htmlBody := fmt.Sprintf("%s has left the ranks of Megajon :(", unsubscriberEmail)
	textBody := "This text body is from megajon.com"

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(Sender),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(htmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(textBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(Sender),
	}

	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	fmt.Println("Email Sent to address: " + Sender)
	fmt.Println(result)
	emailObject := s.OutgoingEmail{
		Sender:    Sender,
		Recipient: unsubscriberEmail,
		HtmlBody:  htmlBody,
		TextBody:  textBody,
		CharSet:   CharSet,
	}
	fmt.Printf("email object: %d", emailObject)
	return emailObject

}
