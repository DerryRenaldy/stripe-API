package repository

import (
	"context"
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
	"stripe-project/models/api/requests"
	"stripe-project/models/api/responses"
	"stripe-project/models/web/responseWeb"
	"stripe-project/repository/helper"
)

func (c *Client) DuplicateValidation(ctx context.Context, req requests.CustomerRequest) ([]responses.Validator, error) {
	// ========== Defining Variables ==========
	var dataValidator []responses.Validator
	// ========== Prepare Query ==========
	querySearch := `SELECT c.phone_number, c.email FROM stripe.customers c 
                    WHERE phone_number=? OR email=?;`

	rows, err := c.DB.QueryContext(ctx, querySearch, req.PhoneNumber, req.Email)
	if err != nil {
		log.Println("ERROR REPOSITORY DUPLICATE VALIDATOR:", err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		res := responses.Validator{}
		err := rows.Scan(&res.PhoneNumber, &res.Email)
		if err != nil {
			log.Println("ERROR REPOSITORY SCAN DATA:", err)
			return nil, err
		}
		dataValidator = append(dataValidator, res)
	}

	return dataValidator, nil
}

func (c *Client) InsertCustomer(ctx context.Context, resAPI *responseWeb.APICustomerResponse) (*responses.CustomerResponse, error) {
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

	defer func(query *sql.Stmt) {
		err := query.Close()
		if err != nil {

		}
	}(query)

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

func (c *Client) InsertCard(ctx context.Context, resAPI *responseWeb.APICardResponse) (*responses.CardResponse, error) {
	// ========== Declaring Variable ==========
	var result responses.CardResponse
	var customer responses.CustomerInfo

	// ========== Prepare Query ==========
	queryInsert := `INSERT INTO stripe.cards (card_id, customer_id, brand) VALUES (?, ?, ?);`
	query, err := c.DB.PrepareContext(ctx, queryInsert)
	if err != nil {
		log.Println("ERROR REPOSITORY PREPARE:", err)
		return nil, err
	}

	defer func(query *sql.Stmt) {
		err := query.Close()
		if err != nil {

		}
	}(query)

	// ========== Exec Query ==========
	_, err = query.ExecContext(ctx, resAPI.CardId, resAPI.CustomerId, resAPI.Brand)
	if err != nil {
		log.Println("ERROR REPOSITORY EXEC:", err)
		return nil, err
	}

	querySearch := `SELECT c.name, c.phone_number FROM stripe.customers c 
                    WHERE c.customer_id = ?;`

	err = c.DB.QueryRowContext(ctx, querySearch, resAPI.CustomerId).Scan(&customer.Name, &customer.PhoneNumber)
	if err != nil {
		log.Println("ERROR REPOSITORY SEARCH INSERT CARD:", err)
		return nil, err
	}

	result = responses.CardResponse{
		CardId:              resAPI.CardId,
		Brand:               resAPI.Brand,
		CustomerName:        customer.Name,
		CustomerPhoneNumber: customer.PhoneNumber,
	}

	return &result, nil
}

func (c *Client) GetCustomerById(ctx context.Context, customerId string) (*responses.CustomerResponse, error) {
	// ========== Declaring Variable ==========
	var result responses.CustomerResponse

	// ========== Query Context ==========
	querySearch := `SELECT c.customer_id, c.name, c.phone_number, c.email, c.status 
					FROM stripe.customers c WHERE customer_id =?;`

	rows, err := c.DB.QueryContext(ctx, querySearch, customerId)
	if err != nil {
		log.Println("ERROR REPOSITORY GET CUSTOMER BY ID:", err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	if rows.Next() {
		err = rows.Scan(&result.CustomerId, &result.Name, &result.PhoneNumber, &result.Email, &result.Status)
		if err != nil {
			log.Println("ERROR REPOSITORY SCAN DATA:", err)
			return nil, err
		}
	} else {
		return nil, errors.New("No Customers Data Found")
	}

	return &result, nil
}

func (c *Client) GetCards(ctx context.Context, brand string, customerId string) ([]responses.GetCardsResponse, error) {
	// ========== Defining Variables ==========
	var result []responses.GetCardsResponse

	querySearch := `SELECT cr.card_id, cr.brand, cr.customer_id FROM stripe.cards cr WHERE 1`

	// Logic Parameter
	if brand == "" && customerId == "" {
		querySearch += `;`

		result, err := helper.QueryEmptyParameter(c.DB, ctx, querySearch, result)
		if err != nil {
			log.Println("ERROR REPOSITORY QUERY EMPTY PARAMETER:", err)
			return nil, err
		}

		return result, nil
	}

	if brand != "" && customerId != "" {
		querySearch += ` AND brand=? && customer_id=?;`

		result, err := helper.QueryAllParameter(c.DB, ctx, querySearch, result, brand, customerId)
		if err != nil {
			log.Println("ERROR REPOSITORY QUERY EMPTY PARAMETER:", err)
			return nil, err
		}

		return result, nil
	}

	if brand != "" {
		querySearch += ` AND brand=?;`
		result, err := helper.QueryBrand(c.DB, ctx, querySearch, result, brand)
		if err != nil {
			log.Println("ERROR REPOSITORY QUERY EMPTY PARAMETER:", err)
			return nil, err
		}

		return result, nil
	}

	if customerId != "" {
		querySearch += ` AND customer_id=?`

		result, err := helper.QueryCustomerId(c.DB, ctx, querySearch, result, customerId)
		if err != nil {
			log.Println("ERROR REPOSITORY QUERY EMPTY PARAMETER:", err)
			return nil, err
		}

		return result, nil
	}

	return nil, errors.New("invalid parameter")
}
