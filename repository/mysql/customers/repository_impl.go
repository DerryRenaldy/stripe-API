package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (c *Client) ChargesValidation(customerId string) (*responses.ValidatorCharges, error) {
	var validation responses.ValidatorCharges
	querySearch := `SELECT c.customer_id, cs.description FROM stripe.customers c 
    				JOIN stripe.customers_status cs ON c.status=cs.status 
                    WHERE c.customer_id=?;`

	rows, err := c.DB.Query(querySearch, customerId)
	if err != nil {
		log.Println("ERROR REPOSITORY VALIDATE CHARGES:", err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	if rows.Next() {
		err = rows.Scan(&validation.CustomerId, &validation.Status)
		if err != nil {
			log.Println("ERROR REPOSITORY SCANNING:", err)
			return nil, err
		}
	}

	return &validation, err
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
	querySearch := `SELECT c.customer_id, c.name, c.phone_number, c.email, cs.description
					FROM stripe.customers_status cs JOIN stripe.customers c ON c.status = cs.status
					WHERE customer_id=?;`

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
	//var customer responses.CustomerInfo

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

	querySearch := `SELECT cr.customer_id, cr.card_id, cr.brand, c.name, c.phone_number FROM stripe.cards cr
					JOIN stripe.customers c ON cr.customer_id=c.customer_id 
                	WHERE cr.card_id=?;`

	err = c.DB.QueryRowContext(ctx, querySearch, resAPI.CardId).Scan(&result.CustomerId, &result.CardId, &result.Brand, &result.CustomerName, &result.CustomerPhoneNumber)
	if err != nil {
		log.Println("ERROR REPOSITORY SEARCH INSERT CARD:", err)
		return nil, err
	}

	// ===== Insert Card Expire Date =====
	expireDate := fmt.Sprintf("%v-%v", resAPI.ExpireMonth, resAPI.ExpireYear)
	result.ExpireDate = expireDate

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

	querySearch := `SELECT cr.card_id, cr.brand, cr.customer_id, c.name, c.phone_number, cs.description
					FROM stripe.customers c JOIN stripe.cards cr ON cr.customer_id = c.customer_id 
					JOIN stripe.customers_status cs ON cs.status=c.status WHERE 1`

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
		querySearch += ` AND cr.brand=? && c.customer_id=? ORDER BY cs.description AND c.name;`

		result, err := helper.QueryAllParameter(c.DB, ctx, querySearch, result, brand, customerId)
		if err != nil {
			log.Println("ERROR REPOSITORY QUERY EMPTY PARAMETER:", err)
			return nil, err
		}

		return result, nil
	}

	if brand != "" {
		querySearch += ` AND cr.brand=? ORDER BY cs.description AND c.name;`
		result, err := helper.QueryBrand(c.DB, ctx, querySearch, result, brand)
		if err != nil {
			log.Println("ERROR REPOSITORY QUERY EMPTY PARAMETER:", err)
			return nil, err
		}

		return result, nil
	}

	if customerId != "" {
		querySearch += ` AND c.customer_id=? ORDER BY cs.description AND c.name;`

		result, err := helper.QueryCustomerId(c.DB, ctx, querySearch, result, customerId)
		if err != nil {
			log.Println("ERROR REPOSITORY QUERY EMPTY PARAMETER:", err)
			return nil, err
		}

		return result, nil
	}

	return nil, errors.New("invalid parameter")
}

func (c *Client) CreateCharges(ctx context.Context, req requests.ChargesRequest, resAPI *responseWeb.APIChargesResponse, customerId string) (*responses.ChargesResponse, error) {
	var result responses.ChargesResponse

	queryInsert := `INSERT INTO charges (payment_id, customer_id, card_id, amount, recipient_url, descriptions) VALUES (?,?,?,?,?,?);`

	query, err := c.DB.PrepareContext(ctx, queryInsert)
	if err != nil {
		log.Println("ERROR REPOSITORY INSERT CHARGES:", err)
		return nil, err
	}

	defer func(query *sql.Stmt) {
		err := query.Close()
		if err != nil {

		}
	}(query)

	_, err = query.ExecContext(ctx, resAPI.PaymentId, customerId, req.CardId, resAPI.Amount, resAPI.ReceiptURL, req.Description)
	if err != nil {
		log.Println("ERROR REPOSITORY EXEC:", err)
		return nil, err
	}

	querySearch := `SELECT ch.payment_id, ch.customer_id, ch.card_id, ch.amount, ch.recipient_url, ch.descriptions, c.name, c.phone_number, cs.description, ch.created_at
					FROM stripe.charges ch JOIN stripe.customers c ON ch.customer_id = c.customer_id 
					JOIN stripe.customers_status cs ON c.status = cs.status WHERE ch.payment_id=?;`

	rows, err := c.DB.QueryContext(ctx, querySearch, resAPI.PaymentId)
	if err != nil {
		log.Println("ERROR REPOSITORY CHARGE SEARCH:", err)
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&result.PaymentId, &result.CustomerId, &result.CardId, &result.Amount, &result.RecipientURL, &result.Descriptions, &result.CustomerName, &result.CustomerPhoneNumber, &result.Status, &result.CreatedAt)
		if err != nil {
			log.Println("ERROR REPOSITORY CHARGE SCAN:", err)
			return nil, err
		}
	}

	return &result, nil
}
