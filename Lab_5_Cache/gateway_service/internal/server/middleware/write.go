package middleware

import (
	"Gateway/internal/server/response_error"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, content any) *response_error.Error {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		log.Println(err)
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
