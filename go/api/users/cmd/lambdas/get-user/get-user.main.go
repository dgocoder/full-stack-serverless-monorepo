package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgocoder/full-stack-serverless-monorepo/go/api/users/internal/controllers"
	"github.com/dgocoder/full-stack-serverless-monorepo/go/pkg/awsapigw"
)

func handler(
	ctx context.Context,
	request *events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	ctrl, err := controllers.NewUserController(ctx)
	if err != nil {
		return awsapigw.SendError(awsapigw.StatusServiceUnavailable, err.Error())
	}

	userID, err := awsapigw.GetParamPath("id", request)
	if err != nil {
		return awsapigw.SendError(awsapigw.StatusBadRequest, err.Error())
	}

	user, err := ctrl.GetUser(ctx, userID)
	if err != nil {
		return awsapigw.SendError(awsapigw.StatusBadRequest, err.Error())
	}

	return awsapigw.SendResponse(awsapigw.StatusOK, user)
}

func main() {
	lambda.Start(handler)
}
