package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgocoder/full-stack-serverless-monorepo/go/api/users/internal/controllers"
)

func handler(
	ctx context.Context,
	request *events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	ctrl, err := controllers.NewUserController()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	rawParam1, found := request.PathParameters["id"]
	if !found {
		return events.APIGatewayProxyResponse{}, errors.New("user id not specified")
	}

	userID, err := url.QueryUnescape(rawParam1)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errors.New("invalid user id specified")
	}

	user, err := ctrl.GetUser(ctx, userID)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	body, err := json.Marshal(user)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
