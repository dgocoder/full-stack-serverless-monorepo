package ddbrepo

import (
	"context"

	"github.com/dgocoder/full-stack-serverless-monorepo/services/users/internal/repositories/types"
)

func (r *ddb) Create(ctx context.Context, user types.CreateUser) (*types.User, error) {
	return nil, nil
}
