package responseWeb

type APICustomerResponse struct {
	CustomerId  string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
	Status      int    `json:"status"`
}

type APICardResponse struct {
	CardId      string `json:"id"`
	Brand       string `json:"brand"`
	CustomerId  string `json:"customer"`
	ExpireMonth int    `json:"exp_month"`
	ExpireYear  int    `json:"exp_year"`
}

type APIChargesResponse struct {
	PaymentId  string `json:"id"`
	ReceiptURL string `json:"receipt_url"`
	Status     string `json:"status"`
	Amount     int    `json:"amount"`
}
