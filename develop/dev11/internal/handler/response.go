package handler

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

type resultResponse struct {
	Result string `json:"result"`
}

func newErrorResponse(w http.ResponseWriter, error string, statusCode int) {
	w.WriteHeader(statusCode)
	result, _ := json.Marshal(errorResponse{error})
	_, _ = w.Write(result)
}

func newResultResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	result, _ := json.Marshal(resultResponse{message})
	_, _ = w.Write(result)
}
