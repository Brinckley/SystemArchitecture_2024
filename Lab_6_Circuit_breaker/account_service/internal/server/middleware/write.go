package middleware

import (
	"account_service/internal/server/response_error"
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, content any) *response_error.Error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, "cannot encode error")
	}
	return nil
}

func WriteWithTokenHeader(w http.ResponseWriter, status int) *response_error.Error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode("Token successfully put to the header")
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, "cannot encode error")
	}
	return nil
}

func WriteWithIdHeader(w http.ResponseWriter, status int) *response_error.Error {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode("Id successfully put to the header")
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, "cannot encode error")
	}
	return nil
}
