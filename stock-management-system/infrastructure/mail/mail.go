package mail

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendMail(message, recipient, subject, sender string, isTest bool) error {

	svc, err := NewSESHandler(isTest)
	if err != nil {
		return err
	}

	_, err = svc.VerifyEmailAddress(&ses.VerifyEmailAddressInput{EmailAddress: aws.String(recipient)})
	if checkAwsError(err) != nil {
		return err
	}

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(message),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	_, err = svc.SendEmail(input)
	return checkAwsError(err)
}

func ValidEmailAddress(email string) bool {
	return strings.HasSuffix(email, "@u-aizu.ac.jp")
}

func checkAwsError(err error) error {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return aerr
		} else {
			return err
		}
	}
	return nil
}
