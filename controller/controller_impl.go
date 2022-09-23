package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	helper.PrintError(err)

	// Fill In the Status value
	apiResponse.Status = customerRequest.Status
	fmt.Println(apiResponse.Status)
	fmt.Printf("%T \n", apiResponse.Status)
	fmt.Println(customerRequest.Status)
	fmt.Printf("%T \n", customerRequest.Status)

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

	// payload is in json format
	payload, err := io.ReadAll(response.Body)
	helper.PrintError(err)

	err = json.Unmarshal(payload, &apiResponse)
	if err != nil {
		log.Println("ERROR UNMARSHAL:", err)
	}

	customerResponse, duplicateMessage, err := c.Services.CreateCustomer(r.Context(), &apiResponse)

	if err != nil {
		err := helper.PrintResponseToJson(w, duplicateMessage)
		if err != nil {
			log.Println("ERROR:", err)
		}
	} else {
		err := helper.PrintResponseToJson(w, customerResponse)
		helper.PrintError(err)
	}
}
