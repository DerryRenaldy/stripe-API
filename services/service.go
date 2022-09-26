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
	CreateCard(ctx context.Context, resAPI *responseWeb.APICardResponse) (*responses.CardResponse, error)
	DuplicateValidation(ctx context.Context, req requests.CustomerRequest) ([]responses.Validator, error)
	GetCustomerById(ctx context.Context, customerId string) (*responses.CustomerResponse, error)
	GetCards(ctx context.Context, brand string, customerId string) ([]responses.GetCardsResponse, error)
}

func NewService(repository repository.Repository) *Services {
	return &Services{
		Repository: repository,
	}
}
