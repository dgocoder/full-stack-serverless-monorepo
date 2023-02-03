package controllers

import (
	"context"

	"github.com/dgocoder/full-stack-serverless-monorepo/go/api/users/internal/repositories"
	"github.com/dgocoder/full-stack-serverless-monorepo/go/api/users/internal/repositories/ddbrepo"
)

type UserController struct {
	repos repositories.RepositorySet
}

// NewUserController to interact with user service actions.
func NewUserController() (*UserController, error) {
	userRepo, err := ddbrepo.NewDDBUserRepository(context.TODO())
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
