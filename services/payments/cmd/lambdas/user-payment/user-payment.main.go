package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.SQSEvent) (error) {
	isLocal := os.Getenv("IS_LOCAL")
	if isLocal == "true" {
		ctx = context.Background()
	}
	for _, v := range event.Records {
		fmt.Println(v.Body)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
