package helper

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type FailResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ReturnSuccessResponse[T any](w http.ResponseWriter, data T, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := SuccessResponse{
		Status: "success",
		Data:   data,
	}
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func ReturnFailResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := FailResponse{
		Status:  "fail",
		Message: message,
	}
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
