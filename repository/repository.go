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
