package mail

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func NewSESHandler(isTest bool) (*ses.SES, error) {
	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}
	if isTest {
		config.Endpoint = aws.String("http://localhost:4579")
	}
	sess, err := session.NewSession(config)

	svc := ses.New(sess)
	if err != nil {
		return svc, err
	}
	return svc, nil
}
