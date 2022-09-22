package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"stripe-project/helper"
	"stripe-project/models/api/requests"
	"stripe-project/models/web/responseWeb"
)

const (
	Content = "application/x-www-form-urlencoded"
	BaseURL = "https://api.stripe.com/v1/customers"
)

func (c *Controller) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customerRequest = requests.CustomerRequest{}
	var client http.Client

	// request json into struct golang
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&customerRequest)
	helper.PrintError(err)

	data := url.Values{}
	data.Add("name", customerRequest.Name)
	data.Add("phone", customerRequest.PhoneNumber)
	data.Add("email", customerRequest.Email)
	dataReader := bytes.NewBufferString(data.Encode())

	request, err := http.NewRequest(http.MethodPost, BaseURL, dataReader)
	helper.PrintError(err)

	request.Header.Set("Authorization", ApiKey)
	request.Header.Set("Content-Type", Content)

	client = http.Client{}

	response, err := client.Do(request)
	helper.PrintError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	// payload berupa response json
	payload, err := io.ReadAll(response.Body)
	helper.PrintError(err)

	var apiResponse responseWeb.APIResponse
	err = json.Unmarshal(payload, &apiResponse)
	helper.PrintError(err)

	customerResponse, err := c.Services.CreateCustomer(r.Context(), &apiResponse)
	helper.PrintError(err)

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(customerResponse)
	helper.PrintError(err)
}
