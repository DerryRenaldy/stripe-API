package requests

type CustomerRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
	Status      int    `json:"status"`
}

type CardRequest struct {
	CardToken string `json:"card_token"`
}

type ChargesRequest struct {
	CardId      string `json:"card_id"`
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
}
