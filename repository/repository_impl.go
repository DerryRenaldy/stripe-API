package repository

import (
	"context"
	log "github.com/sirupsen/logrus"
	"stripe-project/models/api/responses"
	"stripe-project/models/web/responseWeb"
)

func (c *Client) InsertCustomer(ctx context.Context, resAPI *responseWeb.APIResponse) (*responses.CustomerResponse, error) {
	// ========== Declaring Variable ==========
	var result responses.CustomerResponse

	// ========== Logrus Formatter ===========
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// ========== Prepare Query ==========
	queryInsert := `INSERT INTO stripe.customers (customer_id, name, phone_number, email, status) 
					VALUES (?, ?, ?, ?, ?);`
	query, err := c.DB.PrepareContext(ctx, queryInsert)
	if err != nil {
		log.Println("ERROR REPOSITORY PREPARE:", err)
		return nil, err
	}

	defer query.Close()

	// ========== Execute Query ==========
	_, err = query.ExecContext(ctx, resAPI.CustomerId, resAPI.Name, resAPI.PhoneNumber, resAPI.Email, resAPI.Status)
	if err != nil {
		log.Println("ERROR REPOSITORY EXEC:", err)
		return nil, err
	}

	// ========== Search Query ==========
	querySearch := `SELECT c.customer_id, c.name, c.phone_number, c.email, c.status
					FROM stripe.customers c WHERE customer_id=?;`
	err = c.DB.QueryRowContext(ctx, querySearch, resAPI.CustomerId).Scan(&result.CustomerId, &result.Name, &result.PhoneNumber, &result.Email, &result.Status)
	if err != nil {
		log.Println("ERROR REPOSITORY QUERY:", err)
		return nil, err
	}

	return &result, nil
}
