package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sky0621/aws-describe-prj/structure"
)

func NewSqs(sess *session.Session) *sqs.SQS {
	return sqs.New(sess)
}

func GetSqsInformation(cli *sqs.SQS) (*structure.SqsInformation, error) {
	output, err := cli.ListQueues(&sqs.ListQueuesInput{})
	if err != nil {
		return nil, err
	}

	return &structure.SqsInformation{QueueURLs: output.QueueUrls}, nil
}
