package controllers

import (
	"context"

	"github.com/dgocoder/full-stack-serverless-monorepo/services/users/internal/repositories"
	"github.com/dgocoder/full-stack-serverless-monorepo/services/users/internal/repositories/ddbrepo"
)

type UserController struct {
	repos repositories.RepositorySet
}

// NewUserController to interact with user service actions.
func NewUserController(ctx context.Context) (*UserController, error) {
	userRepo, err := ddbrepo.NewDDBUserRepository(ctx)
	if err != nil {
		return nil, err
	}

	repoSet := repositories.RepositorySet{
		Users: userRepo,
	}

	return &UserController{
		repos: repoSet,
	}, nil
}
