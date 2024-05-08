package middleware

import (
	"account_service/internal/server/response_error"
	"fmt"
	"log"
	"net/http"
)

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
		log.Printf("Header %s", w.Header().Values("Auth-token"))
	}
}
