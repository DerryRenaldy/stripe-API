package responses

import "time"

type CustomerResponse struct {
	CustomerId  string `json:"customer_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Status      int    `json:"status"`
}

type CardResponse struct {
	CardId              string `json:"card_id"`
	Brand               string `json:"brand"`
	CustomerName        string `json:"customer_name"`
	CustomerPhoneNumber string `json:"customer_phone_number"`
}

type Validator struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type CustomerInfo struct {
	Name        string
	PhoneNumber string
}

type GetCardsResponse struct {
	CardId     string `json:"card_id"`
	Brand      string `json:"brand"`
	CustomerId string `json:"customer_id"`
}

type PaymentResponse struct {
	CustomerId   string    `json:"customer_id"`
	CardId       string    `json:"card_id"`
	RecipientURL string    `json:"recipient_url"`
	CreatedAt    time.Time `json:"created_at"`
}
