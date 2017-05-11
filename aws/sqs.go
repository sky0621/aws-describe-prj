package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sky0621/aws-describe-prj/config"
	"github.com/sky0621/aws-describe-prj/structure"
)

func NewSqs(sess *session.Session) *sqs.SQS {
	return sqs.New(sess)
}

func GetSqsInformation(cli *sqs.SQS, conf *config.SqsConfig) (*structure.SqsInformation, error) {
	out, err := cli.ListQueues(&sqs.ListQueuesInput{})
	if err != nil {
		return nil, err
	}

	for _, qURL := range out.QueueUrls {
		fmt.Println(qURL)
	}
	fmt.Println("===============================")
	fmt.Println(out.String())
	fmt.Println("===============================")
	fmt.Println(out.GoString())
	fmt.Println("===============================")
	return &structure.SqsInformation{}, nil
}
