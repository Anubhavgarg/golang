package main

import (
	"encoding/json"
)

func ResponseCreation(errorMessage *string, isError bool) *response {
	finalResponse := &response{
		Message: *errorMessage,
		isError: isError,
	}
	return finalResponse
}

func MarshallUnmarshallCreationResponse(apiResponse interface{}, isError bool) *response {
	res1B, _ := json.Marshal(apiResponse)
	//fmt.Println(string(res1B),22)
	res1 := &response{}
	res1.isError = isError
	json.Unmarshal([]byte(string(res1B)), res1)
	return res1
}