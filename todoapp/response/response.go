package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool       `json:"success"`
	Data    any        `json:"data,omitempty"`
	Error   *ErrorInfo `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func OK(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	resp, err := json.Marshal(Response{
		Success: true,
		Data:    data,
	})
	if err != nil {
		panic(err)
	}

	w.Write(resp)
}

func Fail(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	resp, err := json.Marshal(Response{
		Success: false,
		Error:   &ErrorInfo{Code: status, Message: message},
	})
	if err != nil {
		panic(err)
	}

	w.Write(resp)
}
