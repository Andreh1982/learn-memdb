package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// @see https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
func newSession(input Input) (*session.Session, error) {
	return session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Region:                        aws.String(input.AwsRegion),
				Endpoint:                      aws.String(input.AwsEndpoint),
				CredentialsChainVerboseErrors: aws.Bool(true),
			},
			Profile: input.AwsProfile,
		},
	)
}
