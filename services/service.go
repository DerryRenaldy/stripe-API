package services

import (
	"context"
	"stripe-project/models/api/requests"
	"stripe-project/models/api/responses"
	"stripe-project/models/web/responseWeb"
	"stripe-project/repository"
)

type Services struct {
	Repository repository.Repository
}

type Repository interface {
	CreateCustomer(ctx context.Context, resApi *responseWeb.APICustomerResponse) (*responses.CustomerResponse, error)
	CreateCard(ctx context.Context, resAPI *responseWeb.APICardResponse, cusID responseWeb.APICustomerResponse) (*responses.CardResponse, error)
	DuplicateValidation(ctx context.Context, req requests.CustomerRequest) ([]responses.Validator, error)
}

func NewService(repository repository.Repository) *Services {
	return &Services{
		Repository: repository,
	}
}
