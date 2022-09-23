package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"stripe-project/helper"
	"stripe-project/models/api/requests"
	"stripe-project/models/web/responseWeb"
)

const (
	ApiKey  = "Bearer sk_test_51LjzGAG2dQj3VjR6jvsXIEoeLN3Zz8zqWWggrGAwML87OBiiNacGf0qnmQYIhyxy0EXCnfR2v7QVkXIT7bGwdCna00xAXdmzAu"
	Content = "application/x-www-form-urlencoded"
	BaseURL = "https://api.stripe.com/v1/customers"
)

func (c *Controller) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	// ========== Define Variable ==========
	var customerRequest = requests.CustomerRequest{}
	var client http.Client
	var apiResponse responseWeb.APIResponse

	// ========== Controller Logic ==========
	// request json into struct golang
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&customerRequest)
	if err != nil {
		log.Println("ERROR UNMARSHAL:", err)
		helper.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	// Fill In the Status value
	apiResponse.Status = customerRequest.Status

	data := url.Values{}
	data.Add("name", customerRequest.Name)
	data.Add("phone", customerRequest.PhoneNumber)
	data.Add("email", customerRequest.Email)
	dataReader := bytes.NewBufferString(data.Encode())

	request, err := http.NewRequest(http.MethodPost, BaseURL, dataReader)
	if err != nil {
		log.Println("ERROR CREATE NEW REQUEST:", err)
		helper.RespondWithError(w, http.StatusBadGateway, err.Error())
	}

	request.Header.Set("Authorization", ApiKey)
	request.Header.Set("Content-Type", Content)

	client = http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("ERROR EXECUTE REQUEST:", err)
		helper.RespondWithError(w, http.StatusBadGateway, err.Error())
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
	}

	err = json.Unmarshal(payload, &apiResponse)
	if err != nil {
		log.Println("ERROR UNMARSHAL:", err)
		helper.RespondWithError(w, http.StatusBadRequest, "Bad Request")
	}

	customerResponse, err := c.Services.CreateCustomer(r.Context(), &apiResponse)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Duplicate Email or Phone Number")
	} else {
		helper.RespondWithJSON(w, http.StatusOK, customerResponse)
	}
}
