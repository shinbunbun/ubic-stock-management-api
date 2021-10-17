package mail

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func NewSESHandler(isTest bool) *ses.SES {
	config := &aws.Config{
		Region: aws.String("ap-northeast-1"),
	}
	if isTest {
		config.Endpoint = aws.String("http://localhost:4579")
	}
	sess, _ := session.NewSession(config)

	svc := ses.New(sess)
	return svc
}
