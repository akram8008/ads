package domain

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func (a *ApiResponse) Respond(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(a.Code)

	resp, _ := json.Marshal(a)

	w.Write(resp)
}

func (a *ApiResponse) Set(code int, message string, payload interface{}) {
	a.Code = code
	a.Message = message
	a.Payload = payload
}
