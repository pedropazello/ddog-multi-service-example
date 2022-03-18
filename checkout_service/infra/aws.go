package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func AwsSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func NewSQSClient() *sqs.SQS {
	return sqs.New(AwsSession(), Config())
}

func Config() *aws.Config {
	return &aws.Config{
		Endpoint:    aws.String("http://localstack:4566"),
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("foo", "var", ""),
	}
}
