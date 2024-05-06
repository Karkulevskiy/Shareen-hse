package lib

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Type    string          `json:"type"`
	Status  int             `json:"status"`
	Payload json.RawMessage `json:"payload"`
}

func HTPPErr(w http.ResponseWriter, eventType string, status int) {

	type ResponseErr struct {
		Status int `json:"status"`
	}

	statusData, _ := json.Marshal(ResponseErr{
		Status: status,
	})

	respErr := Response{
		Type:    eventType,
		Status:  status,
		Payload: statusData,
	}

	resp, _ := json.Marshal(respErr)
	http.Error(w, "", status)
	w.Write(resp)
}
