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
