package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/dgocoder/full-stack-serverless-monorepo/pkg/go/awsconf"
)

func handler(ctx context.Context, event events.DynamoDBEvent) (error) {
	isLocal := os.Getenv("IS_LOCAL")
	if isLocal == "true" {
		ctx = context.TODO()
	}

	topicArn := os.Getenv("USER_CREATED_TOPIC")
	
	client, err := awsconf.NewSNSConf(ctx)

	if err != nil {
		return err
	}

	for _, v := range event.Records {
		by, err := json.Marshal(v.Change)
		if err != nil {
			return err
		}
		m := string(by)
		_, err = client.Publish(ctx, &sns.PublishInput{
			Message: &m,
			TopicArn: &topicArn,
		})
		if err != nil {
			return err
		}
		fmt.Println(string(by))
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
