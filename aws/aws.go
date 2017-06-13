package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSession() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		// TODO configから読むようにする
		Region: aws.String(endpoints.ApNortheast1RegionID),
	}))
}
