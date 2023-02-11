package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgocoder/full-stack-serverless-monorepo/pkg/go/awsapigw"
	"github.com/dgocoder/full-stack-serverless-monorepo/services/users/internal/controllers"
)

func handler(
	ctx context.Context,
	request *events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	isLocal := os.Getenv("IS_LOCAL")
	if isLocal == "true" {
		ctx = context.TODO()
	}
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
