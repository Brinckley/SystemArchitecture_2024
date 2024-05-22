package middleware

import (
	"Gateway/internal/server/response_error"
	"encoding/json"
	"io"
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

func WriteJsonFromResponse(w http.ResponseWriter, status int, r *http.Response) *response_error.Error {
	w.WriteHeader(status)
	_, err := io.Copy(w, r.Body)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, "cannot copy body")
	}
	return nil
}
