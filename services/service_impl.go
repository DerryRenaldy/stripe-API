package services

import (
	"context"
	log "github.com/sirupsen/logrus"
	"stripe-project/models/api/requests"
	"stripe-project/models/api/responses"
	"stripe-project/models/web/responseWeb"
)

func (s *Services) CreateCustomer(ctx context.Context, resAPI *responseWeb.APICustomerResponse) (*responses.CustomerResponse, error) {
	// ========== Service Logic ==========
	resp, err := s.Repository.InsertCustomer(ctx, resAPI)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		return nil, err
	}

	return resp, nil
}

func (s *Services) CreateCard(ctx context.Context, resAPI *responseWeb.APICardResponse) (*responses.CardResponse, error) {
	resp, err := s.Repository.InsertCard(ctx, resAPI)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		return nil, err
	}

	return resp, nil
}

func (s *Services) DuplicateValidation(ctx context.Context, req requests.CustomerRequest) ([]responses.Validator, error) {
	validator, err := s.Repository.DuplicateValidation(ctx, req)
	if err != nil {
		log.Println("ERROR SERVICE DUPLICATE VALIDATION:", err)
		return nil, err
	}
	return validator, err
}
