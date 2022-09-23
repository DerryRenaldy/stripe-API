package services

import (
	"context"
	"stripe-project/models/api/responses"
	"stripe-project/models/web/responseWeb"
	"stripe-project/repository"
)

type Services struct {
	Repository repository.Repository
}

type Repository interface {
	CreateCustomer(ctx context.Context, resApi *responseWeb.APIResponse) (*responses.CustomerResponse, error)
}

func NewService(repository repository.Repository) *Services {
	return &Services{
		Repository: repository,
	}
}
