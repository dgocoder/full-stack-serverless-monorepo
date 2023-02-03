package ddbrepo

import (
	"context"

	"github.com/dgocoder/full-stack-serverless-monorepo/go/api/users/internal/repositories/types"
)

func (r *ddb) Get(ctx context.Context, id string) (*types.User, error) {
	user := types.User{
		ID:    id,
		Email: "test@test.com",
	}

	return &user, nil
}
