package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"stripe-project/controller"
	"stripe-project/helper"
	"stripe-project/repository"
	"stripe-project/services"
)

type Server struct {
	sql        repository.Repository
	services   services.Repository
	controller controller.Repository
}

func Register() *Server {
	SVR := &Server{}

	SVR.sql = repository.NewClient()

	SVR.services = services.NewService(SVR.sql)

	SVR.controller = controller.NewController(SVR.services)

	return SVR
}

func (s *Server) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/createCustomer", s.controller.CreateCustomer).Methods(http.MethodPost)
	r.HandleFunc("/createCard/{id}", s.controller.CreateCard).Methods(http.MethodPost)

	fmt.Println("Server Listening At Port 8010...")
	err := http.ListenAndServe(":8010", r)
	helper.PrintError(err)
}
