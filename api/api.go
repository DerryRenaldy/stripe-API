package api

//func HitAPI(w http.ResponseWriter, req interface{}) {
//
//	data := url.Values{}
//	data.Add("name", req.Name)
//	data.Add("phone", req.PhoneNumber)
//	data.Add("email", req.Email)
//	dataReader := bytes.NewBufferString(data.Encode())
//
//	request, err := http.NewRequest(http.MethodPost, BaseURL+"/v1/customers", dataReader)
//	if err != nil {
//		log.Println("ERROR CREATE NEW REQUEST:", err)
//		helper.RespondWithError(w, http.StatusBadGateway, err.Error())
//		return
//	}
//
//	request.Header.Set("Authorization", ApiKey)
//	request.Header.Set("Content-Type", Content)
//
//	client = http.Client{}
//
//	response, err := client.Do(request)
//	if err != nil {
//		log.Println("ERROR EXECUTE REQUEST:", err)
//		helper.RespondWithError(w, http.StatusBadGateway, err.Error())
//		return
//	}
//
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//
//		}
//	}(response.Body)
//
//	// payload is in json format
//	payload, err := io.ReadAll(response.Body)
//	if err != nil {
//		log.Println("ERROR PARSING PAYLOAD:", err)
//		helper.RespondWithError(w, http.StatusExpectationFailed, "parsing failed")
//		return
//	}
//
//	err = json.Unmarshal(payload, &apiResponse)
//	if err != nil {
//		log.Println("ERROR UNMARSHAL:", err)
//		helper.RespondWithError(w, http.StatusBadRequest, "Bad Request")
//		return
//	}
//}
