package services

import (
	"context"
	"stripe-project/helper"
	"stripe-project/models/api/responses"
	"stripe-project/models/web/responseWeb"
)

func (s *Services) CreateCustomer(ctx context.Context, resAPI *responseWeb.APIResponse) (*responses.CustomerResponse, error) {
	resp, err := s.Repository.InsertCustomer(ctx, resAPI)
	helper.PrintError(err)

	return resp, nil
}
