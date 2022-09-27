package helper

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"stripe-project/models/api/responses"
)

func Scan(rows *sql.Rows, result []responses.GetCardsResponse) ([]responses.GetCardsResponse, error) {
	for rows.Next() {
		row := responses.GetCardsResponse{}
		err := rows.Scan(&row.CardId, &row.Brand, &row.CustomerId, &row.CustomerName, &row.CustomerPhoneNumber, &row.Status)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}
	return result, nil
}

func QueryEmptyParameter(db *sql.DB, ctx context.Context, query string, result []responses.GetCardsResponse) ([]responses.GetCardsResponse, error) {
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Println("ERROR REPOSITORY GET CARDS:", err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	result, err = Scan(rows, result)
	if err != nil {
		log.Println("ERROR REPOSITORY SCAN DATA:", err)
		return nil, err
	}

	return result, nil
}

func QueryBrand(db *sql.DB, ctx context.Context, query string, result []responses.GetCardsResponse, brand string) ([]responses.GetCardsResponse, error) {
	rows, err := db.QueryContext(ctx, query, brand)
	if err != nil {
		log.Println("ERROR REPOSITORY GET CARDS:", err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	result, err = Scan(rows, result)
	if err != nil {
		log.Println("ERROR REPOSITORY SCAN DATA:", err)
		return nil, err
	}

	return result, nil
}

func QueryCustomerId(db *sql.DB, ctx context.Context, query string, result []responses.GetCardsResponse, CustomerId string) ([]responses.GetCardsResponse, error) {
	rows, err := db.QueryContext(ctx, query, CustomerId)
	if err != nil {
		log.Println("ERROR REPOSITORY GET CARDS:", err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	result, err = Scan(rows, result)
	if err != nil {
		log.Println("ERROR REPOSITORY SCAN DATA:", err)
		return nil, err
	}

	return result, nil
}

func QueryAllParameter(db *sql.DB, ctx context.Context, query string, result []responses.GetCardsResponse, brand string, customerId string) ([]responses.GetCardsResponse, error) {
	rows, err := db.QueryContext(ctx, query, brand, customerId)
	if err != nil {
		log.Println("ERROR REPOSITORY GET CARDS:", err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	result, err = Scan(rows, result)
	if err != nil {
		log.Println("ERROR REPOSITORY SCAN DATA:", err)
		return nil, err
	}

	return result, nil
}
