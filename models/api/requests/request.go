package requests

type CustomerRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
}
