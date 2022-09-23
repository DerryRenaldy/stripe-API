package services

import (
	"context"
	log "github.com/sirupsen/logrus"
	"stripe-project/models/api/responses"
	"stripe-project/models/web/responseWeb"
)

func (s *Services) CreateCustomer(ctx context.Context, resAPI *responseWeb.APIResponse) (*responses.CustomerResponse, error) {
	// ========== Service Logic ==========
	resp, err := s.Repository.InsertCustomer(ctx, resAPI)
	if err != nil {
		log.Println("ERROR SERVICE:", err)
		return nil, err
	}

	return resp, nil
}
