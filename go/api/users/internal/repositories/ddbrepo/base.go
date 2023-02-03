package ddbrepo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/dgocoder/full-stack-serverless-monorepo/go/api/users/internal/repositories"
	"github.com/dgocoder/full-stack-serverless-monorepo/go/pkg/awsconf"
)

type ddb struct {
	client *dynamodb.Client
}

// NewDDBUserRepository dynamodb repository for user objects.
func NewDDBUserRepository(ctx context.Context) (repositories.UserRepository, error) {
	ddbClient, err := awsconf.NewBaseDynamoDBConf(ctx)
	if err != nil {
		return nil, err
	}

	return &ddb{
		client: ddbClient,
	}, nil
}
