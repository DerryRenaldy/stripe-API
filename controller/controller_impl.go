package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"stripe-project/helper"
	"stripe-project/models/api/requests"
	"stripe-project/models/web/responseWeb"
)

const (
	ApiKey  = "Bearer sk_test_51LjzGAG2dQj3VjR6jvsXIEoeLN3Zz8zqWWggrGAwML87OBiiNacGf0qnmQYIhyxy0EXCnfR2v7QVkXIT7bGwdCna00xAXdmzAu"
	Content = "application/x-www-form-urlencoded"
	BaseURL = "https://api.stripe.com"
)

func (c *Controller) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	// ========== Define Variable ==========
	var customerRequest = requests.CustomerRequest{}
	var client http.Client
	var apiResponse responseWeb.APICustomerResponse
	var errValidate error

	// ========== Controller Logic ==========
	// request json into struct golang
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&customerRequest)
	if err != nil {
		log.Println("ERROR UNMARSHAL:", err)
		helper.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	// ===== Validate Duplicate Card =====
	validator, errValidate := c.Services.DuplicateValidation(r.Context(), customerRequest)
	if validator == nil && errValidate == nil {
		// Fill In the Status value
		apiResponse.Status = customerRequest.Status

		data := url.Values{}
		data.Add("name", customerRequest.Name)
		data.Add("phone", customerRequest.PhoneNumber)
		data.Add("email", customerRequest.Email)
		dataReader := bytes.NewBufferString(data.Encode())

		request, err := http.NewRequest(http.MethodPost, BaseURL+"/v1/customers", dataReader)
		if err != nil {
			log.Println("ERROR CREATE NEW REQUEST:", err)
			helper.RespondWithError(w, http.StatusBadGateway, err.Error())
			return
		}

		request.Header.Set("Authorization", ApiKey)
		request.Header.Set("Content-Type", Content)

		client = http.Client{}

		response, err := client.Do(request)
		if err != nil {
			log.Println("ERROR EXECUTE REQUEST:", err)
			helper.RespondWithError(w, http.StatusBadGateway, err.Error())
			return
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(response.Body)

		// payload is in json format
		payload, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println("ERROR PARSING PAYLOAD:", err)
			helper.RespondWithError(w, http.StatusExpectationFailed, "parsing failed")
			return
		}

		err = json.Unmarshal(payload, &apiResponse)
		if err != nil {
			log.Println("ERROR UNMARSHAL:", err)
			helper.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		customerResponse, err := c.Services.CreateCustomer(r.Context(), &apiResponse)

		if err != nil {
			helper.RespondWithError(w, http.StatusBadRequest, "Duplicate Email or Phone Number")
			return
		} else {
			helper.RespondWithJSON(w, http.StatusOK, customerResponse)
		}
	} else {
		helper.RespondWithError(w, http.StatusBadRequest, "Duplicate Email or Phone Number")
		return
	}

}

func (c *Controller) CreateCard(w http.ResponseWriter, r *http.Request) {
	// ========== Define Variable ==========
	var cardRequest = requests.CardRequest{}
	var client http.Client
	var apiResponse responseWeb.APICardResponse

	// ========== Get Customer Id ==========
	params := mux.Vars(r)
	id := params["id"]

	// ========== Controller Logic ==========
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cardRequest)
	if err != nil {
		log.Println("ERROR UNMARSHAL:", err)
		helper.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	data := url.Values{}
	data.Add("source", cardRequest.CardToken)
	dataReader := bytes.NewBufferString(data.Encode())

	request, err := http.NewRequest(http.MethodPost, BaseURL+"/v1/customers/"+id+"/sources", dataReader)
	if err != nil {
		log.Println("ERROR CREATE NEW REQUEST:", err)
		helper.RespondWithError(w, http.StatusBadGateway, err.Error())
		return
	}

	request.Header.Set("Authorization", ApiKey)
	request.Header.Set("Content-Type", Content)

	client = http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("ERROR EXECUTE REQUEST:", err)
		helper.RespondWithError(w, http.StatusBadGateway, err.Error())
		return
	}

	if response.StatusCode != http.StatusOK {
		helper.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	// payload is in json format
	payload, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("ERROR PARSING PAYLOAD:", err)
		helper.RespondWithError(w, http.StatusExpectationFailed, "parsing failed")
		return
	}

	err = json.Unmarshal(payload, &apiResponse)
	if err != nil {
		log.Println("ERROR UNMARSHAL:", err)
		helper.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	cardResponse, err := c.Services.CreateCard(r.Context(), &apiResponse)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		helper.RespondWithJSON(w, http.StatusOK, cardResponse)
	}
}

func (c *Controller) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId := r.URL.Query().Get("customerId")

	customerResponse, err := c.Services.GetCustomerById(r.Context(), customerId)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, customerResponse)
}

func (c *Controller) GetCards(w http.ResponseWriter, r *http.Request) {
	customerId := r.URL.Query().Get("customerId")
	brand := r.URL.Query().Get("brand")

	cardResponse, err := c.Services.GetCards(r.Context(), brand, customerId)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, cardResponse)
}

func (c *Controller) CreateCharges(w http.ResponseWriter, r *http.Request) {
	// ========== Declaring Variables ==========
	var chargeRequest requests.ChargesRequest
	var client http.Client
	var apiResponse responseWeb.APIChargesResponse
	var errValidate error

	// Get Customer id From URL Parameter
	customerId := r.URL.Query().Get("customerId")

	// ========== Controller Logic ==========
	// request json into struct golang
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&chargeRequest)
	if err != nil {
		log.Println("ERROR UNMARSHAL:", err)
		helper.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	// ========== Validate Status customer ==========
	validate, errValidate := c.Services.ChargesValidation(customerId)

	if validate != nil && errValidate == nil && validate.Status == "active" {

		data := url.Values{}
		data.Add("amount", strconv.Itoa(chargeRequest.Amount))
		data.Add("currency", chargeRequest.Currency)
		data.Add("description", chargeRequest.Description)
		data.Add("customer", customerId)
		dataReader := bytes.NewBufferString(data.Encode())

		request, err := http.NewRequest(http.MethodPost, BaseURL+"/v1/charges", dataReader)
		if err != nil {
			log.Println("ERROR CREATE NEW REQUEST:", err)
			helper.RespondWithError(w, http.StatusBadGateway, err.Error())
			return
		}

		request.Header.Set("Authorization", ApiKey)
		request.Header.Set("Content-Type", Content)

		client = http.Client{}

		response, err := client.Do(request)
		if err != nil {
			log.Println("ERROR EXECUTE REQUEST:", err)
			helper.RespondWithError(w, http.StatusBadGateway, err.Error())
			return
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(response.Body)

		// payload is in json format
		payload, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println("ERROR PARSING PAYLOAD:", err)
			helper.RespondWithError(w, http.StatusExpectationFailed, "parsing failed")
			return
		}

		err = json.Unmarshal(payload, &apiResponse)
		if err != nil {
			log.Println("ERROR UNMARSHAL:", err)
			helper.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		customerResponse, err := c.Services.CreateCharges(r.Context(), chargeRequest, &apiResponse, customerId)

		if err != nil {
			helper.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		} else {
			helper.RespondWithJSON(w, http.StatusOK, customerResponse)
		}
	} else {
		helper.RespondWithError(w, http.StatusBadRequest, "Customer is Not Active")
		return
	}
}
