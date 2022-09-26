package controller

import (
	"net/http"
	"stripe-project/services"
)

type Controller struct {
	Services services.Repository
}

type Repository interface {
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	CreateCard(w http.ResponseWriter, r *http.Request)
	GetCustomerById(w http.ResponseWriter, r *http.Request)
	GetCards(w http.ResponseWriter, r *http.Request)
}

func NewController(services services.Repository) *Controller {
	return &Controller{Services: services}
}
