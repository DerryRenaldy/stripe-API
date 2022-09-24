package responses

type CustomerResponse struct {
	CustomerId  string  `json:"id"`
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone"`
	Email       string  `json:"email"`
	Status      int     `json:"status"`
	CardId      *string `json:"card_id"`
}

type CardResponse struct {
	CardId string `json:"card_id"`
	Brand  string `json:"brand"`
}

type Validator struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
