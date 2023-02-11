package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.DynamoDBEvent) (error) {
	isLocal := os.Getenv("IS_LOCAL")
	if isLocal == "true" {
		ctx = context.TODO()
	}
	for _, v := range event.Records {
		by, err := json.Marshal(v.Change)
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
