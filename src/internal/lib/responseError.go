package lib

import "encoding/json"

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func Err(message string, status int) []byte {
	respErr := ResponseError{message, status}

	data, _ := json.Marshal(respErr)

	return data
}
