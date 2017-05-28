package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sky0621/aws-describe-prj/structure"
)

func NewDynamoDB(sess *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(sess)
}

func GetDynamoDBInformation(cli *dynamodb.DynamoDB) (*structure.DynamoDBInformation, error) {
	output, err := cli.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return nil, err
	}

	return &structure.DynamoDBInformation{TableNames: output.TableNames}, nil
}
