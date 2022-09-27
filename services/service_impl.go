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

func (s *Services) GetCustomerById(ctx context.Context, customerId string) (*responses.CustomerResponse, error) {
	resp, err := s.Repository.GetCustomerById(ctx, customerId)
	if err != nil {
		log.Println("ERROR SERVICE GET CUSTOMER BY ID:", err)
		return nil, err
	}

	return resp, nil
}

func (s *Services) GetCards(ctx context.Context, brand string, customerId string) ([]responses.GetCardsResponse, error) {
	resp, err := s.Repository.GetCards(ctx, brand, customerId)
	if err != nil {
		log.Println("ERROR SERVICE GET CARDS:", err)
		return nil, err
	}

	return resp, nil
}

func (s *Services) CreateCharges(ctx context.Context, req requests.ChargesRequest, resAPI *responseWeb.APIChargesResponse, customerId string) ([]responses.ChargesResponse, error) {
	resp, err := s.Repository.CreateCharges(ctx, req, resAPI, customerId)
	if err != nil {
		log.Println("ERROR SERVICE CREATE CHARGES:", err)
		return nil, err
	}

	return resp, err
}

func (s *Services) ChargesValidation(customerId string) (*responses.ValidatorCharges, error) {
	validate, err := s.Repository.ChargesValidation(customerId)
	if err != nil {
		log.Println("ERROR SERVICE DUPLICATE VALIDATION:", err)
		return nil, err
	}

	return validate, err
}
