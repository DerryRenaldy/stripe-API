package services

import (
	"context"
	log "github.com/sirupsen/logrus"
	"stripe-project/models/api/responses"
	"stripe-project/models/web/responseWeb"
)

func (s *Services) CreateCustomer(ctx context.Context, resAPI *responseWeb.APIResponse) (*responses.CustomerResponse, *responses.DuplicateCustomerResponse, error) {
	// ========== Define Message ==========
	message := &responses.DuplicateCustomerResponse{
		Message: "Phone Number or Email is Already Exist",
	}

	// ========== Service Logic ==========
	resp, err := s.Repository.InsertCustomer(ctx, resAPI)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		return nil, message, err
	}

	return resp, nil, nil
}
