package repositories

import (
	"context"

	"github.com/dgocoder/full-stack-serverless-monorepo/go/api/users/internal/repositories/types"
)

type UserRepository interface {
	Create(context.Context, types.CreateUser) (*types.User, error)
	Get(ctx context.Context, id string) (*types.User, error)
}

type RepositorySet struct {
	Users UserRepository
}
