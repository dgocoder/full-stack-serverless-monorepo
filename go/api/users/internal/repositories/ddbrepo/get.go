package ddbrepo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/dgocoder/full-stack-serverless-monorepo/go/api/users/internal/repositories/types"
)

func (r *ddb) Get(ctx context.Context, id string) (*types.User, error) {
	key := fmt.Sprintf("USER#%s", id)

	res, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		Key: map[string]ddbtypes.AttributeValue{
			"pk": &ddbtypes.AttributeValueMemberS{Value: key},
			"sk": &ddbtypes.AttributeValueMemberS{Value: key},
		},
		TableName: &r.tableName,
	})
	if err != nil {
		return nil, err
	}

	usr := types.User{}

	err = attributevalue.UnmarshalMap(res.Item, &usr)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}
