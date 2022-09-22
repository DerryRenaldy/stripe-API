package repository

import (
	"context"
	"stripe-project/helper"
	"stripe-project/models/api/responses"
	"stripe-project/models/web/responseWeb"
)

func (c *Client) InsertCustomer(ctx context.Context, resAPI *responseWeb.APIResponse) (*responses.CustomerResponse, error) {
	queryInsert := `INSERT INTO stripe.customers (customer_id, name, phone_number, email) 
					VALUES (?, ?, ?, ?);`
	query, err := c.DB.PrepareContext(ctx, queryInsert)
	helper.PrintError(err)

	defer query.Close()

	_, err = query.ExecContext(ctx, resAPI.CustomerId, resAPI.Name, resAPI.PhoneNumber, resAPI.Email)
	helper.PrintError(err)

	querySearch := `SELECT c.customer_id, c.name, c.phone_number, c.email 
					FROM stripe.customers c WHERE customer_id=?;`

	var result responses.CustomerResponse
	err = c.DB.QueryRowContext(ctx, querySearch, resAPI.CustomerId).Scan(&result.CustomerId, &result.Name, &result.PhoneNumber, &result.Email)
	helper.PrintError(err)

	return &result, nil
}
