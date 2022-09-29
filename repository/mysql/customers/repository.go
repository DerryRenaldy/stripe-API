package repository

import (
	"context"
	"database/sql"
	"log"
	"stripe-project/database"
	"stripe-project/models/api/requests"
	"stripe-project/models/api/responses"
	"stripe-project/models/web/responseWeb"
)

type Client struct {
	DB *sql.DB
}

type Repository interface {
	DuplicateValidation(ctx context.Context, req requests.CustomerRequest) ([]responses.Validator, error)
	InsertCustomer(ctx context.Context, resAPI *responseWeb.APICustomerResponse) (*responses.CustomerResponse, error)
	InsertCard(ctx context.Context, resAPI *responseWeb.APICardResponse) (*responses.CardResponse, error)
	GetCustomerById(ctx context.Context, customerId string) (*responses.CustomerResponse, error)
	GetCards(ctx context.Context, brand string, customerId string) ([]responses.GetCardsResponse, error)
	CreateCharges(ctx context.Context, req requests.ChargesRequest, resAPI *responseWeb.APIChargesResponse, customerId string) (*responses.ChargesResponse, error)
	ChargesValidation(customerId string) (*responses.ValidatorCharges, error)
}

func NewClient() *Client {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Client{
		DB: db,
	}
}