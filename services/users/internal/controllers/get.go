package controllers

import (
	"context"

	"github.com/dgocoder/full-stack-serverless-monorepo/services/users/internal/repositories/types"
)

// GetUser retrieves the user with given ID.
func (u *UserController) GetUser(ctx context.Context, id string) (*types.User, error) {
	user, err := u.repos.Users.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
