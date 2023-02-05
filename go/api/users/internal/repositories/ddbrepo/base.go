package ddbrepo

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/dgocoder/full-stack-serverless-monorepo/go/api/users/internal/repositories"
	"github.com/dgocoder/full-stack-serverless-monorepo/go/pkg/awsconf"
)

type ddb struct {
	client    *dynamodb.Client
	tableName string
}

// NewDDBUserRepository dynamodb repository for user objects.
func NewDDBUserRepository(ctx context.Context) (repositories.UserRepository, error) {
	ddbClient, err := awsconf.NewBaseDynamoDBConf(ctx)
	if err != nil {
		return nil, err
	}

	tableName := os.Getenv("USERS_TABLE_NAME")

	return &ddb{
		client:    ddbClient,
		tableName: tableName,
	}, nil
}
