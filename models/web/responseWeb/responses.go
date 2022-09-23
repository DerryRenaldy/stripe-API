package responseWeb

type APICustomerResponse struct {
	CustomerId  string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
	Status      int    `json:"status"`
}

type APICardResponse struct {
	CardId string `json:"id"`
	Brand  string `json:"brand"`
}
