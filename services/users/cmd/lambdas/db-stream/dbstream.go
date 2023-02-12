package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgocoder/full-stack-serverless-monorepo/pkg/go/awsconf"
	"github.com/dgocoder/full-stack-serverless-monorepo/services/users/internal/repositories/types"
)

func handler(ctx context.Context, event events.DynamoDBEvent) error {
	isLocal := os.Getenv("IS_LOCAL")
	if isLocal == "true" {
		ctx = context.TODO()
	}

	user := awsconf.NewStreamEntity(awsconf.StreamEntityInput{
		EventSource:  "USER",
		DomainEntity: types.User{},
		EventTypes:   []awsconf.DDBStreamEventType{awsconf.DDBStreamEventTypeInsert},
	})

	err := awsconf.DDBEntityToSNS(ctx, event, []awsconf.StreamEntity{user})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
