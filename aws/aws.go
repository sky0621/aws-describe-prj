package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSession() (*session.Session, error) {
	// Credentialは環境変数セット済の前提
	awsCfg := &aws.Config{}
	awsCfg.Credentials = credentials.NewEnvCredentials()

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		return nil, err
	}

	return sess, nil
}
