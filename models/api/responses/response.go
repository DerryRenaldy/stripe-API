package responses

import "time"

type CustomerResponse struct {
	CustomerId  string `json:"customer_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Status      string `json:"status"`
}

type CardResponse struct {
	CustomerId          string `json:"customer_id"`
	CardId              string `json:"card_id"`
	Brand               string `json:"brand"`
	ExpireDate          string `json:"expire_date"`
	CustomerName        string `json:"customer_name"`
	CustomerPhoneNumber string `json:"customer_phone_number"`
}

type Validator struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type ValidatorCharges struct {
	CustomerId string `json:"customer_id"`
	Status     string `json:"status"`
}

type CustomerInfo struct {
	Name        string
	PhoneNumber string
}

type GetCardsResponse struct {
	CardId              string `json:"card_id"`
	Brand               string `json:"brand"`
	CustomerId          string `json:"customer_id"`
	CustomerName        string `json:"customer_name"`
	CustomerPhoneNumber string `json:"customer_phone_number"`
	Status              string `json:"status"`
}

type ChargesResponse struct {
	PaymentId           string    `json:"payment_id"`
	CustomerId          string    `json:"customer_id"`
	CardId              string    `json:"card_id"`
	Amount              int       `json:"amount"`
	RecipientURL        string    `json:"recipient_url"`
	Descriptions        string    `json:"descriptions"`
	CustomerName        string    `json:"customer_name"`
	CustomerPhoneNumber string    `json:"customer_phone_number"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"created_at"`
}
