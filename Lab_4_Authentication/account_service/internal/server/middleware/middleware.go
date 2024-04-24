package middleware

import (
	"account_service/internal/server/response_error"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func WriteJson(w http.ResponseWriter, status int, content any) *response_error.Error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Length", strconv.Itoa(1<<20))
	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, "cannot encode error")
	}
	return nil
}

func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("account service received request method: %s, header: %s, body: %s\n", r.Method, r.Header, r.Body)
		next.ServeHTTP(w, r)
	})
}

type apiFunc func(w http.ResponseWriter, r *http.Request) *response_error.Error

func MakeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			err := WriteJson(w, err.StatusCode(), err.Message())
			if err != nil {
				log.Println(fmt.Errorf("unable to write error data error : %s", err))
				return
			}
		}
	}
}
