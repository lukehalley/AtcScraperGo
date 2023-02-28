package aws

import (
	"atcscraper/src/env"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func NewAWSClient() *lambda.Lambda {

	AWSSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	AWSRegion := env.LoadEnv("AWS_DEFAULT_REGION")

	AWSClient := lambda.New(AWSSession, &aws.Config{Region: aws.String(AWSRegion)})

	return AWSClient

}
