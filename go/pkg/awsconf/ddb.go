package awsconf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewBaseDynamoDBConf(ctx context.Context) (*dynamodb.Client, error) {
	conf, err := NewAWSConfig(ctx)
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(conf), nil
}
