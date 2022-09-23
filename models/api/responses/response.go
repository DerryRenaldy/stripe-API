package responses

type CustomerResponse struct {
	CustomerId  string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
	Status      int    `json:"status"`
}

type DuplicateCustomerResponse struct {
	Message string `json:"message"`
}
