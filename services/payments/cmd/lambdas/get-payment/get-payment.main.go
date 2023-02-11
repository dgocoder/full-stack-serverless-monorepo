package main

import (
	"fmt"
	"math/rand"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("WOAH2222 NEW")
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, from payments service %v", fmt.Sprint(rand.Intn(100))),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
